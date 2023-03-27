// SPDX-License-Identifier: MIT

pragma solidity ^0.8.9;

contract Messages {
    string public message;
    
    constructor(string memory initialMessage) {
        message = initialMessage;
    }
    
    function setNewMessage(string memory newMessage) public {
        message = newMessage;
    }
}