const path = require("path");
const fs = require("fs");
const solc = require("solc");

const messagesPath = path.resolve(__dirname, "Messages.sol");
const source = fs.readFileSync(messagesPath, "utf8");

const input = {
  language: "Solidity",
  sources: {
    "Messages.sol": {
      content: source,
    },
  },
  settings: {
    outputSelection: {
      "*": {
        "*": ["*"],
      },
    },
  },
};

module.exports = JSON.parse(solc.compile(JSON.stringify(input))).contracts[
  "Messages.sol"
].Messages;
