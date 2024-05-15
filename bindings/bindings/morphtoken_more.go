// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"morph-l2/bindings/solc"
)

const MorphTokenStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/l2/system/MorphToken.sol:MorphToken\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/l2/system/MorphToken.sol:MorphToken\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/l2/system/MorphToken.sol:MorphToken\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1016_storage\"},{\"astId\":1003,\"contract\":\"contracts/l2/system/MorphToken.sol:MorphToken\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/l2/system/MorphToken.sol:MorphToken\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1015_storage\"},{\"astId\":1005,\"contract\":\"contracts/l2/system/MorphToken.sol:MorphToken\",\"label\":\"_name\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_string_storage\"},{\"astId\":1006,\"contract\":\"contracts/l2/system/MorphToken.sol:MorphToken\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_string_storage\"},{\"astId\":1007,\"contract\":\"contracts/l2/system/MorphToken.sol:MorphToken\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_uint256\"},{\"astId\":1008,\"contract\":\"contracts/l2/system/MorphToken.sol:MorphToken\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1009,\"contract\":\"contracts/l2/system/MorphToken.sol:MorphToken\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1010,\"contract\":\"contracts/l2/system/MorphToken.sol:MorphToken\",\"label\":\"_epochInflationRates\",\"offset\":0,\"slot\":\"106\",\"type\":\"t_array(t_struct(EpochInflationRate)1017_storage)dyn_storage\"},{\"astId\":1011,\"contract\":\"contracts/l2/system/MorphToken.sol:MorphToken\",\"label\":\"_inflations\",\"offset\":0,\"slot\":\"107\",\"type\":\"t_mapping(t_uint256,t_uint256)\"},{\"astId\":1012,\"contract\":\"contracts/l2/system/MorphToken.sol:MorphToken\",\"label\":\"_inflationMintedEpochs\",\"offset\":0,\"slot\":\"108\",\"type\":\"t_uint256\"},{\"astId\":1013,\"contract\":\"contracts/l2/system/MorphToken.sol:MorphToken\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"109\",\"type\":\"t_array(t_uint256)1014_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_struct(EpochInflationRate)1017_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct IMorphToken.EpochInflationRate[]\",\"numberOfBytes\":\"32\"},\"t_array(t_uint256)1014_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[38]\",\"numberOfBytes\":\"1216\"},\"t_array(t_uint256)1015_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1016_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_struct(EpochInflationRate)1017_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IMorphToken.EpochInflationRate\",\"numberOfBytes\":\"64\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var MorphTokenStorageLayout = new(solc.StorageLayout)

var MorphTokenDeployedBin = "0x608060405234801561000f575f80fd5b506004361061018f575f3560e01c8063715018a6116100dd578063a457c2d711610088578063cd4281d011610063578063cd4281d0146103af578063dd62ed3e146103d6578063f2fde38b1461041b575f80fd5b8063a457c2d714610381578063a9059cbb14610394578063c553f7b3146103a7575f80fd5b8063944fa746116100b8578063944fa7461461034757806395d89b4114610366578063a29bfb2c1461036e575f80fd5b8063715018a6146102fa578063807de443146103025780638da5cb5b14610329575f80fd5b8063395093511161013d5780636d0c4a26116101185780636d0c4a26146102845780636fe0e395146102b257806370a08231146102c5575f80fd5b806339509351146102105780633d9353fe14610223578063405abb411461026f575f80fd5b806318160ddd1161016d57806318160ddd146101e657806323b872dd146101ee578063313ce56714610201575f80fd5b806306fdde0314610193578063095ea7b3146101b15780630b88a984146101d4575b5f80fd5b61019b61042e565b6040516101a891906115f1565b60405180910390f35b6101c46101bf366004611683565b6104be565b60405190151581526020016101a8565b606c545b6040519081526020016101a8565b6067546101d8565b6101c46101fc3660046116ab565b6104d7565b604051601281526020016101a8565b6101c461021e366004611683565b6104fa565b61024a7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a8565b61028261027d3660046116e4565b610545565b005b610297610292366004611704565b6106b5565b604080518251815260209283015192810192909252016101a8565b6102826102c03660046117ef565b61070c565b6101d86102d336600461185e565b73ffffffffffffffffffffffffffffffffffffffff165f9081526068602052604090205490565b61028261095f565b61024a7f000000000000000000000000000000000000000000000000000000000000000081565b60335473ffffffffffffffffffffffffffffffffffffffff1661024a565b6101d8610355366004611704565b5f908152606b602052604090205490565b61019b610972565b61028261037c366004611704565b610981565b6101c461038f366004611683565b610d1c565b6101c46103a2366004611683565b610dc6565b606a546101d8565b61024a7f000000000000000000000000000000000000000000000000000000000000000081565b6101d86103e436600461187e565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260696020908152604080832093909416825291909152205490565b61028261042936600461185e565b610dd3565b60606065805461043d906118af565b80601f0160208091040260200160405190810160405280929190818152602001828054610469906118af565b80156104b45780601f1061048b576101008083540402835291602001916104b4565b820191905f5260205f20905b81548152906001019060200180831161049757829003601f168201915b5050505050905090565b5f336104cb818585610e8a565b60019150505b92915050565b5f336104e4858285610ff1565b6104ef8585856110c7565b506001949350505050565b335f81815260696020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906104cb908290869061054090879061192d565b610e8a565b61054d6112ca565b606a805461055d90600190611940565b8154811061056d5761056d611953565b905f5260205f209060020201600101548111610610576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f6566666563746976652065706f636873206166746572206d757374206265206760448201527f726561746572207468616e206265666f7265000000000000000000000000000060648201526084015b60405180910390fd5b60408051808201825283815260208101838152606a80546001810182555f91825292517f116fea137db6e131133e7f2bab296045d8f41cc5607279db17b218cab0929a5160029094029384015590517f116fea137db6e131133e7f2bab296045d8f41cc5607279db17b218cab0929a52909201919091559051829184917fbe80a5653ecb34691beafb0fb70004d50f9032b798f68a2f73a137c4f98ab3f49190a35050565b604080518082019091525f8082526020820152606a82815481106106db576106db611953565b905f5260205f2090600202016040518060400160405290815f82015481526020016001820154815250509050919050565b5f54610100900460ff161580801561072a57505f54600160ff909116105b806107435750303b15801561074357505f5460ff166001145b6107cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610607565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561082b575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61083361134b565b606561083f86826119c9565b50606661084c85826119c9565b5061085733846113e9565b6040805180820182528381525f60208201818152606a805460018101825590835292517f116fea137db6e131133e7f2bab296045d8f41cc5607279db17b218cab0929a51600290940293840155517f116fea137db6e131133e7f2bab296045d8f41cc5607279db17b218cab0929a5290920191909155905183907fbe80a5653ecb34691beafb0fb70004d50f9032b798f68a2f73a137c4f98ab3f4908390a38015610958575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b6109676112ca565b6109705f6114dc565b565b60606066805461043d906118af565b337f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1614610a20576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f6f6e6c79207265636f726420636f6e747261637420616c6c6f776564000000006044820152606401610607565b807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663766718086040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a8a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610aae9190611ae5565b11610b3b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f746865207370656369666965642074696d6520686173206e6f7420796574206260448201527f65656e20726561636865640000000000000000000000000000000000000000006064820152608401610607565b606c54811015610ba7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f616c6c20696e666c6174696f6e73206d696e74656400000000000000000000006044820152606401610607565b606c545b818111610d0a575f606a5f81548110610bc657610bc6611953565b5f9182526020822060029091020154606a54909250610be790600190611940565b90505b8015610c505782606a8281548110610c0457610c04611953565b905f5260205f2090600202016001015411610c3e57606a8181548110610c2c57610c2c611953565b905f5260205f2090600202015f015491505b80610c4881611afc565b915050610bea565b50662386f26fc1000081606754610c679190611b30565b610c719190611b47565b5f838152606b60205260409020819055610cac907f0000000000000000000000000000000000000000000000000000000000000000906113e9565b817f0d82c0920038b8dc7f633e18585f37092ba957b84876fcf833d6841f69eaa327606b5f8581526020019081526020015f2054604051610cef91815260200190565b60405180910390a25080610d0281611b7f565b915050610bab565b50610d1681600161192d565b606c5550565b335f81815260696020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610db9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f64656372656173656420616c6c6f77616e63652062656c6f77207a65726f00006044820152606401610607565b6104ef8286868403610e8a565b5f336104cb8185856110c7565b610ddb6112ca565b73ffffffffffffffffffffffffffffffffffffffff8116610e7e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610607565b610e87816114dc565b50565b73ffffffffffffffffffffffffffffffffffffffff8316610f07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f617070726f76652066726f6d20746865207a65726f20616464726573730000006044820152606401610607565b73ffffffffffffffffffffffffffffffffffffffff8216610f84576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f617070726f766520746f20746865207a65726f206164647265737300000000006044820152606401610607565b73ffffffffffffffffffffffffffffffffffffffff8381165f8181526069602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152606960209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146110c157818110156110b4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f696e73756666696369656e7420616c6c6f77616e6365000000000000000000006044820152606401610607565b6110c18484848403610e8a565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316611144576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f7472616e736665722066726f6d20746865207a65726f206164647265737300006044820152606401610607565b73ffffffffffffffffffffffffffffffffffffffff82166111c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f7472616e7366657220746f20746865207a65726f2061646472657373000000006044820152606401610607565b73ffffffffffffffffffffffffffffffffffffffff83165f9081526068602052604090205481811015611250576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f7472616e7366657220616d6f756e7420657863656564732062616c616e6365006044820152606401610607565b73ffffffffffffffffffffffffffffffffffffffff8085165f8181526068602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906112bc9086815260200190565b60405180910390a350505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610970576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610607565b5f54610100900460ff166113e1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610607565b610970611552565b73ffffffffffffffffffffffffffffffffffffffff8216611466576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6d696e7420746f20746865207a65726f206164647265737300000000000000006044820152606401610607565b8060675f828254611477919061192d565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f818152606860209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff166115e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610607565b610970336114dc565b5f602080835283518060208501525f5b8181101561161d57858101830151858201604001528201611601565b505f6040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461167e575f80fd5b919050565b5f8060408385031215611694575f80fd5b61169d8361165b565b946020939093013593505050565b5f805f606084860312156116bd575f80fd5b6116c68461165b565b92506116d46020850161165b565b9150604084013590509250925092565b5f80604083850312156116f5575f80fd5b50508035926020909101359150565b5f60208284031215611714575f80fd5b5035919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112611757575f80fd5b813567ffffffffffffffff808211156117725761177261171b565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156117b8576117b861171b565b816040528381528660208588010111156117d0575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f805f8060808587031215611802575f80fd5b843567ffffffffffffffff80821115611819575f80fd5b61182588838901611748565b9550602087013591508082111561183a575f80fd5b5061184787828801611748565b949794965050505060408301359260600135919050565b5f6020828403121561186e575f80fd5b6118778261165b565b9392505050565b5f806040838503121561188f575f80fd5b6118988361165b565b91506118a66020840161165b565b90509250929050565b600181811c908216806118c357607f821691505b6020821081036118fa577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808201808211156104d1576104d1611900565b818103818111156104d1576104d1611900565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b601f8211156119c457805f5260205f20601f840160051c810160208510156119a55750805b601f840160051c820191505b81811015610958575f81556001016119b1565b505050565b815167ffffffffffffffff8111156119e3576119e361171b565b6119f7816119f184546118af565b84611980565b602080601f831160018114611a49575f8415611a135750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611add565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611a9557888601518255948401946001909101908401611a76565b5085821015611ad157878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f60208284031215611af5575f80fd5b5051919050565b5f81611b0a57611b0a611900565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b80820281158282048414176104d1576104d1611900565b5f82611b7a577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611baf57611baf611900565b506001019056fea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(MorphTokenStorageLayoutJSON), MorphTokenStorageLayout); err != nil {
		panic(err)
	}

	layouts["MorphToken"] = MorphTokenStorageLayout
	deployedBytecodes["MorphToken"] = MorphTokenDeployedBin
}
