// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { BytesUtils } from "../BytesUtils.sol";
import { RLPReader } from "../rlp/RLPReader.sol";
import { RLPWriter } from "../rlp/RLPWriter.sol";

/**
 * @title MerkleTrie
 * @notice MerkleTrie is a small library for verifying standard Ethereum Merkle-Patricia trie
 *         inclusion proofs. By default, this library assumes a hexary trie. One can change the
 *         trie radix constant to support other trie radixes.
 */
library MerkleTrie {
    /**
     * @notice Struct representing a node in the trie.
     */
    struct TrieNode {
        bytes encoded;
        RLPReader.RLPItem[] decoded;
    }

    /**
     * @notice Determines the number of elements per branch node.
     */
    uint256 internal constant TREE_RADIX = 16;

    /**
     * @notice Branch nodes have TREE_RADIX elements and one value element.
     */
    uint256 internal constant BRANCH_NODE_LENGTH = TREE_RADIX + 1;

    /**
     * @notice Leaf nodes and extension nodes have two elements, a `path` and a `value`.
     */
    uint256 internal constant LEAF_OR_EXTENSION_NODE_LENGTH = 2;

    /**
     * @notice Prefix for even-nibbled extension node paths.
     */
    uint8 internal constant PREFIX_EXTENSION_EVEN = 0;

    /**
     * @notice Prefix for odd-nibbled extension node paths.
     */
    uint8 internal constant PREFIX_EXTENSION_ODD = 1;

    /**
     * @notice Prefix for even-nibbled leaf node paths.
     */
    uint8 internal constant PREFIX_LEAF_EVEN = 2;

    /**
     * @notice Prefix for odd-nibbled leaf node paths.
     */
    uint8 internal constant PREFIX_LEAF_ODD = 3;

    /**
     * @notice RLP representation of `NULL`.
     */
    bytes1 internal constant RLP_NULL = bytes1(0x80);

    /**
     * @notice Verifies a proof that a given key/value pair is present in the trie.
     *
     * @param _key   Key of the node to search for, as a hex string.
     * @param _value Value of the node to search for, as a hex string.
     * @param _proof Merkle trie inclusion proof for the desired node. Unlike traditional Merkle
     *               trees, this proof is executed top-down and consists of a list of RLP-encoded
     *               nodes that make a path down to the target node.
     * @param _root  Known root of the Merkle trie. Used to verify that the included proof is
     *               correctly constructed.
     *
     * @return Whether or not the proof is valid.
     */
    function verifyInclusionProof(
        bytes memory _key,
        bytes memory _value,
        bytes memory _proof,
        bytes32 _root
    ) internal pure returns (bool) {
        (bool exists, bytes memory value) = get(_key, _proof, _root);
        return (exists && BytesUtils.equal(_value, value));
    }

    /**
     * @notice Retrieves the value associated with a given key.
     *
     * @param _key   Key to search for, as hex bytes.
     * @param _proof Merkle trie inclusion proof for the key.
     * @param _root  Known root of the Merkle trie.
     *
     * @return Whether or not the key exists.
     * @return Value of the key if it exists.
     */
    function get(
        bytes memory _key,
        bytes memory _proof,
        bytes32 _root
    ) internal pure returns (bool, bytes memory) {
        TrieNode[] memory proof = _parseProof(_proof);
        (uint256 pathLength, bytes memory keyRemainder, bool isFinalNode) = _walkNodePath(
            proof,
            _key,
            _root
        );

        bool exists = keyRemainder.length == 0;

        require(exists || isFinalNode, "Provided proof is invalid.");

        bytes memory value = exists ? _getNodeValue(proof[pathLength - 1]) : bytes("");

        return (exists, value);
    }

    /**
     * @notice Walks through a proof using a provided key.
     *
     * @param _proof Inclusion proof to walk through.
     * @param _key   Key to use for the walk.
     * @param _root  Known root of the trie.
     *
     * @return Length of the final path
     * @return Portion of the key remaining after the walk.
     * @return Whether or not we've hit a dead end.
     */
    // solhint-disable-next-line code-complexity
    function _walkNodePath(
        TrieNode[] memory _proof,
        bytes memory _key,
        bytes32 _root
    )
        private
        pure
        returns (
            uint256,
            bytes memory,
            bool
        )
    {
        uint256 pathLength = 0;
        bytes memory key = BytesUtils.toNibbles(_key);

        bytes32 currentNodeID = _root;
        uint256 currentKeyIndex = 0;
        uint256 currentKeyIncrement = 0;
        TrieNode memory currentNode;

        // Proof is top-down, so we start at the first element (root).
        for (uint256 i = 0; i < _proof.length; i++) {
            currentNode = _proof[i];
            currentKeyIndex += currentKeyIncrement;

            // Keep track of the proof elements we actually need.
            // It's expensive to resize arrays, so this simply reduces gas costs.
            pathLength += 1;

            if (currentKeyIndex == 0) {
                // First proof element is always the root node.
                require(keccak256(currentNode.encoded) == currentNodeID, "Invalid root hash");
            } else if (currentNode.encoded.length >= 32) {
                // Nodes 32 bytes or larger are hashed inside branch nodes.
                require(
                    keccak256(currentNode.encoded) == currentNodeID,
                    "Invalid large internal hash"
                );
            } else {
                // Nodes smaller than 31 bytes aren't hashed.
                require(
                    bytes32(currentNode.encoded) == currentNodeID,
                    "Invalid internal node hash"
                );
            }

            if (currentNode.decoded.length == BRANCH_NODE_LENGTH) {
                if (currentKeyIndex == key.length) {
                    // We've hit the end of the key
                    // meaning the value should be within this branch node.
                    break;
                } else {
                    // We're not at the end of the key yet.
                    // Figure out what the next node ID should be and continue.
                    uint8 branchKey = uint8(key[currentKeyIndex]);
                    RLPReader.RLPItem memory nextNode = currentNode.decoded[branchKey];
                    currentNodeID = _getNodeID(nextNode);
                    currentKeyIncrement = 1;
                    continue;
                }
            } else if (currentNode.decoded.length == LEAF_OR_EXTENSION_NODE_LENGTH) {
                bytes memory path = _getNodePath(currentNode);
                uint8 prefix = uint8(path[0]);
                uint8 offset = 2 - (prefix % 2);
                bytes memory pathRemainder = BytesUtils.slice(path, offset);
                bytes memory keyRemainder = BytesUtils.slice(key, currentKeyIndex);
                uint256 sharedNibbleLength = _getSharedNibbleLength(pathRemainder, keyRemainder);

                if (prefix == PREFIX_LEAF_EVEN || prefix == PREFIX_LEAF_ODD) {
                    if (
                        pathRemainder.length == sharedNibbleLength &&
                        keyRemainder.length == sharedNibbleLength
                    ) {
                        // The key within this leaf matches our key exactly.
                        // Increment the key index to reflect that we have no remainder.
                        currentKeyIndex += sharedNibbleLength;
                    }

                    // We've hit a leaf node, so our next node should be NULL.
                    currentNodeID = bytes32(RLP_NULL);
                    break;
                } else if (prefix == PREFIX_EXTENSION_EVEN || prefix == PREFIX_EXTENSION_ODD) {
                    if (sharedNibbleLength != pathRemainder.length) {
                        // Our extension node is not identical to the remainder.
                        // We've hit the end of this path
                        // updates will need to modify this extension.
                        currentNodeID = bytes32(RLP_NULL);
                        break;
                    } else {
                        // Our extension shares some nibbles.
                        // Carry on to the next node.
                        currentNodeID = _getNodeID(currentNode.decoded[1]);
                        currentKeyIncrement = sharedNibbleLength;
                        continue;
                    }
                } else {
                    revert("Received a node with an unknown prefix");
                }
            } else {
                revert("Received an unparseable node.");
            }
        }

        // If our node ID is NULL, then we're at a dead end.
        bool isFinalNode = currentNodeID == bytes32(RLP_NULL);
        return (pathLength, BytesUtils.slice(key, currentKeyIndex), isFinalNode);
    }

    /**
     * @notice Parses an RLP-encoded proof into something more useful.
     *
     * @param _proof RLP-encoded proof to parse.
     *
     * @return Proof parsed into easily accessible structs.
     */
    function _parseProof(bytes memory _proof) private pure returns (TrieNode[] memory) {
        RLPReader.RLPItem[] memory nodes = RLPReader.readList(_proof);
        TrieNode[] memory proof = new TrieNode[](nodes.length);

        for (uint256 i = 0; i < nodes.length; i++) {
            bytes memory encoded = RLPReader.readBytes(nodes[i]);
            proof[i] = TrieNode({ encoded: encoded, decoded: RLPReader.readList(encoded) });
        }

        return proof;
    }

    /**
     * @notice Picks out the ID for a node. Node ID is referred to as the "hash" within the
     *         specification, but nodes < 32 bytes are not actually hashed.
     *
     * @param _node Node to pull an ID for.
     *
     * @return ID for the node, depending on the size of its contents.
     */
    function _getNodeID(RLPReader.RLPItem memory _node) private pure returns (bytes32) {
        bytes memory nodeID;

        if (_node.length < 32) {
            // Nodes smaller than 32 bytes are RLP encoded.
            nodeID = RLPReader.readRawBytes(_node);
        } else {
            // Nodes 32 bytes or larger are hashed.
            nodeID = RLPReader.readBytes(_node);
        }

        return bytes32(nodeID);
    }

    /**
     * @notice Gets the path for a leaf or extension node.
     *
     * @param _node Node to get a path for.
     *
     * @return Node path, converted to an array of nibbles.
     */
    function _getNodePath(TrieNode memory _node) private pure returns (bytes memory) {
        return BytesUtils.toNibbles(RLPReader.readBytes(_node.decoded[0]));
    }

    /**
     * @notice Gets the path for a node.
     *
     * @param _node Node to get a value for.
     *
     * @return Node value, as hex bytes.
     */
    function _getNodeValue(TrieNode memory _node) private pure returns (bytes memory) {
        return RLPReader.readBytes(_node.decoded[_node.decoded.length - 1]);
    }

    /**
     * @notice Utility; determines the number of nibbles shared between two nibble arrays.
     *
     * @param _a First nibble array.
     * @param _b Second nibble array.
     *
     * @return Number of shared nibbles.
     */
    function _getSharedNibbleLength(bytes memory _a, bytes memory _b)
        private
        pure
        returns (uint256)
    {
        uint256 i = 0;
        while (_a.length > i && _b.length > i && _a[i] == _b[i]) {
            i++;
        }
        return i;
    }
}
