// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { ERC721 } from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import { IERC165 } from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import { Strings } from "@openzeppelin/contracts/utils/Strings.sol";
import { IOptimismMintableERC721 } from "./IOptimismMintableERC721.sol";

/**
 * @title OptimismMintableERC721
 * @notice This contract is the remote representation for some token that lives on another network,
 *         typically an Optimism representation of an Ethereum-based token. Standard reference
 *         implementation that can be extended or modified according to your needs.
 */
contract OptimismMintableERC721 is ERC721, IOptimismMintableERC721 {
    /**
     * @inheritdoc IOptimismMintableERC721
     */
    address public remoteToken;

    /**
     * @inheritdoc IOptimismMintableERC721
     */
    address public bridge;

    /**
     * @notice Base token URI for this token.
     */
    string public baseTokenURI;

    /**
     * @param _bridge      Address of the bridge on this network.
     * @param _remoteToken Address of the corresponding token on the other network.
     * @param _name        ERC721 name.
     * @param _symbol      ERC721 symbol.
     */
    constructor(
        address _bridge,
        address _remoteToken,
        string memory _name,
        string memory _symbol
    ) ERC721(_name, _symbol) {
        remoteToken = _remoteToken;
        bridge = _bridge;

        // Creates a base URI in the format specified by EIP-681:
        // https://eips.ethereum.org/EIPS/eip-681
        baseTokenURI = string(
            abi.encodePacked(
                "ethereum:",
                Strings.toHexString(uint160(_remoteToken)),
                "@",
                Strings.toString(block.chainid),
                "/tokenURI?uint256="
            )
        );
    }

    /**
     * @notice Modifier that prevents callers other than the bridge from calling the function.
     */
    modifier onlyBridge() {
        require(msg.sender == bridge, "OptimismMintableERC721: only bridge can call this function");
        _;
    }

    /**
     * @inheritdoc IOptimismMintableERC721
     */
    function mint(address _to, uint256 _tokenId) external virtual onlyBridge {
        _mint(_to, _tokenId);

        emit Mint(_to, _tokenId);
    }

    /**
     * @inheritdoc IOptimismMintableERC721
     */
    function burn(address _from, uint256 _tokenId) external virtual onlyBridge {
        _burn(_tokenId);

        emit Burn(_from, _tokenId);
    }

    /**
     * @notice Checks if a given interface ID is supported by this contract.
     *
     * @param _interfaceId The interface ID to check.
     *
     * @return True if the interface ID is supported, false otherwise.
     */
    function supportsInterface(bytes4 _interfaceId)
        public
        view
        override(ERC721, IERC165)
        returns (bool)
    {
        bytes4 iface1 = type(IERC165).interfaceId;
        bytes4 iface2 = type(IOptimismMintableERC721).interfaceId;
        return
            _interfaceId == iface1 ||
            _interfaceId == iface2 ||
            super.supportsInterface(_interfaceId);
    }

    /**
     * @notice Returns the base token URI.
     *
     * @return Base token URI.
     */
    function _baseURI() internal view virtual override returns (string memory) {
        return baseTokenURI;
    }
}
