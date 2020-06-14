pragma solidity >=0.4.16 <0.7.0;

contract TexaResultStorage {
    mapping(string=>string) Latest_CID;

    function LogTexaResultURL(string memory aiName, string memory resultURL) public {
        Latest_CID[aiName] = resultURL;
    }

    function GetGlobalTexaResultURL(string memory aiName) public view returns (string memory globalResultURL, bool status) {
        if (bytes(Latest_CID[aiName]).length > 0) {
            return (string(abi.encodePacked("https://explore.ipld.io/#/explore/", Latest_CID[aiName])), true);
        }
        return ("", false);
    }
    
    function GetLocalTexaResultURL(string memory aiName) public view returns (string memory localResultURL, bool status) {
        if (bytes(Latest_CID[aiName]).length > 0) {
            return (string(abi.encodePacked("http://localhost:3000/#/explore/", Latest_CID[aiName])), true);
        }
        return ("", false);
    }
}
