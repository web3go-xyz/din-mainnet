// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

/// @title IBigStepper
/// @notice An interface for a contract with a state transition function that
///         will accept a pre state and return a post state.
/// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
/// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⠶⢅⠒⢄⢔⣶⡦⣤⡤⠄⣀⠀⠀⠀⠀⠀⠀⠀
/// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠨⡏⠀⠀⠈⠢⣙⢯⣄⠀⢨⠯⡺⡘⢄⠀⠀⠀⠀⠀
/// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣶⡆⠀⠀⠀⠀⠈⠓⠬⡒⠡⣀⢙⡜⡀⠓⠄⠀⠀⠀
/// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡷⠿⣧⣀⡀⠀⠀⠀⠀⠀⠀⠉⠣⣞⠩⠥⠀⠼⢄⠀⠀
/// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡇⠀⠀⠀⠉⢹⣶⠒⠒⠂⠈⠉⠁⠘⡆⠀⣿⣿⠫⡄⠀
/// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⢶⣤⣀⡀⠀⠀⢸⡿⠀⠀⠀⠀⠀⢀⠞⠀⠀⢡⢨⢀⡄⠀
/// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⡒⣿⢿⡤⠝⡣⠉⠁⠚⠛⠀⠤⠤⣄⡰⠁⠀⠀⠀⠉⠙⢸⠀⠀
/// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡤⢯⡌⡿⡇⠘⡷⠀⠁⠀⠀⢀⣰⠢⠲⠛⣈⣸⠦⠤⠶⠴⢬⣐⣊⡂⠀
/// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⡪⡗⢫⠞⠀⠆⣀⠻⠤⠴⠐⠚⣉⢀⠦⠂⠋⠁⠀⠁⠀⠀⠀⠀⢋⠉⠇⠀
/// ⠀⠀⠀⠀⣀⡤⠐⠒⠘⡹⠉⢸⠇⠸⠀⠀⠀⠀⣀⣤⠴⠚⠉⠈⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠼⠀⣾⠀
/// ⠀⠀⠀⡰⠀⠉⠉⠀⠁⠀⠀⠈⢇⠈⠒⠒⠘⠈⢀⢡⡂⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⠀⢸⡄
/// ⠀⠀⠸⣿⣆⠤⢀⡀⠀⠀⠀⠀⢘⡌⠀⠀⣀⣀⣀⡈⣤⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⢸⡇
/// ⠀⠀⢸⣀⠀⠉⠒⠐⠛⠋⠭⠭⠍⠉⠛⠒⠒⠒⠀⠒⠚⠛⠛⠛⠩⠭⠭⠭⠭⠤⠤⠤⠤⠤⠭⠭⠉⠓⡆
/// ⠀⠀⠘⠿⣷⣶⣤⣤⣀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣤⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇
/// ⠀⠀⠀⠀⠀⠉⠙⠛⠛⠻⠿⢿⣿⣿⣷⣶⣶⣶⣤⣤⣀⣁⣛⣃⣒⠿⠿⠿⠤⠠⠄⠤⠤⢤⣛⣓⣂⣻⡇
/// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠉⠉⠙⠛⠻⠿⠿⠿⢿⣿⣿⣿⣷⣶⣶⣾⣿⣿⣿⣿⠿⠟⠁
/// ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠈⠉⠉⠉⠉⠁⠀⠀⠀⠀⠀
interface IBigStepper {
    /// @notice Performs a single instruction step from a given prestate and returns the poststate
    ///         hash.
    /// @param _stateData The preimage of the prestate hash.
    /// @param _proof A proof for the inclusion of the prestate's memory in the merkle tree.
    /// @return postState_ The poststate hash after the instruction step.
    function step(bytes calldata _stateData, bytes calldata _proof)
        external
        returns (bytes32 postState_);

    /// @notice Returns the preimage oracle used by the stepper.
    function oracle() external view returns (IPreimageOracle oracle_);
}

/// @notice Temporary interface for the `IPreimageOracle`. Remove once we've upgraded
///         the cannon contracts to a newer version of solc.
interface IPreimageOracle {
    /// @notice Loads a local data part into the preimage oracle.
    /// @param _partOffset The offset of the local data part.
    /// @param _ident The identifier of the local data.
    /// @param _part The local data part.
    /// @param _size The size of the local data.
    function loadLocalPart(
        uint256 _partOffset,
        uint256 _ident,
        bytes32 _part,
        uint256 _size
    ) external;
}
