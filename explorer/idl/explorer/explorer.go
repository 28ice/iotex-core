// Code generated by idl2go from JSON generated by Barrister v0.1.6
package explorer

import (
	"fmt"
	"github.com/coopernurse/barrister-go"
	"reflect"
)

const BarristerVersion string = "0.1.6"
const BarristerChecksum string = "5bfc05ce00c6db2acc1370679969d159"
const BarristerDateGenerated int64 = 1529472246257000000

type CoinStatistic struct {
	Height    int64 `json:"height"`
	Supply    int64 `json:"supply"`
	Transfers int64 `json:"transfers"`
	Votes     int64 `json:"votes"`
	Aps       int64 `json:"aps"`
}

type BlockGenerator struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Block struct {
	ID         string         `json:"ID"`
	Height     int64          `json:"height"`
	Timestamp  int64          `json:"timestamp"`
	Transfers  int64          `json:"transfers"`
	Votes      int64          `json:"votes"`
	GenerateBy BlockGenerator `json:"generateBy"`
	Amount     int64          `json:"amount"`
	Forged     int64          `json:"forged"`
	Size       int64          `json:"size"`
}

type Transfer struct {
	ID        string `json:"ID"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    int64  `json:"amount"`
	Fee       int64  `json:"fee"`
	Timestamp int64  `json:"timestamp"`
	BlockID   string `json:"blockID"`
}

type AddressDetails struct {
	Address      string `json:"address"`
	TotalBalance int64  `json:"totalBalance"`
	Nonce        int64  `json:"nonce"`
}

type Explorer interface {
	GetAddressBalance(address string) (int64, error)
	GetAddressDetails(address string) (AddressDetails, error)
	GetLastTransfersByRange(startBlockHeight int64, offset int64, limit int64, showCoinBase bool) ([]Transfer, error)
	GetTransferByID(transferID string) (Transfer, error)
	GetTransfersByAddress(address string) ([]Transfer, error)
	GetTransfersByBlockID(blockID string) ([]Transfer, error)
	GetLastBlocksByRange(offset int64, limit int64) ([]Block, error)
	GetBlockByID(blockID string) (Block, error)
	GetCoinStatistic() (CoinStatistic, error)
}

func NewExplorerProxy(c barrister.Client) Explorer {
	return ExplorerProxy{c, barrister.MustParseIdlJson([]byte(IdlJsonRaw))}
}

type ExplorerProxy struct {
	client barrister.Client
	idl    *barrister.Idl
}

func (_p ExplorerProxy) GetAddressBalance(address string) (int64, error) {
	_res, _err := _p.client.Call("Explorer.getAddressBalance", address)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getAddressBalance").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf(int64(0)), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.(int64)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getAddressBalance returned invalid type: %v", _t)
			return int64(0), &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return int64(0), _err
}

func (_p ExplorerProxy) GetAddressDetails(address string) (AddressDetails, error) {
	_res, _err := _p.client.Call("Explorer.getAddressDetails", address)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getAddressDetails").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf(AddressDetails{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.(AddressDetails)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getAddressDetails returned invalid type: %v", _t)
			return AddressDetails{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return AddressDetails{}, _err
}

func (_p ExplorerProxy) GetLastTransfersByRange(startBlockHeight int64, offset int64, limit int64, showCoinBase bool) ([]Transfer, error) {
	_res, _err := _p.client.Call("Explorer.getLastTransfersByRange", startBlockHeight, offset, limit, showCoinBase)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getLastTransfersByRange").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf([]Transfer{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.([]Transfer)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getLastTransfersByRange returned invalid type: %v", _t)
			return []Transfer{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return []Transfer{}, _err
}

func (_p ExplorerProxy) GetTransferByID(transferID string) (Transfer, error) {
	_res, _err := _p.client.Call("Explorer.getTransferByID", transferID)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getTransferByID").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf(Transfer{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.(Transfer)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getTransferByID returned invalid type: %v", _t)
			return Transfer{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return Transfer{}, _err
}

func (_p ExplorerProxy) GetTransfersByAddress(address string) ([]Transfer, error) {
	_res, _err := _p.client.Call("Explorer.getTransfersByAddress", address)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getTransfersByAddress").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf([]Transfer{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.([]Transfer)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getTransfersByAddress returned invalid type: %v", _t)
			return []Transfer{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return []Transfer{}, _err
}

func (_p ExplorerProxy) GetTransfersByBlockID(blockID string) ([]Transfer, error) {
	_res, _err := _p.client.Call("Explorer.getTransfersByBlockID", blockID)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getTransfersByBlockID").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf([]Transfer{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.([]Transfer)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getTransfersByBlockID returned invalid type: %v", _t)
			return []Transfer{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return []Transfer{}, _err
}

func (_p ExplorerProxy) GetLastBlocksByRange(offset int64, limit int64) ([]Block, error) {
	_res, _err := _p.client.Call("Explorer.getLastBlocksByRange", offset, limit)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getLastBlocksByRange").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf([]Block{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.([]Block)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getLastBlocksByRange returned invalid type: %v", _t)
			return []Block{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return []Block{}, _err
}

func (_p ExplorerProxy) GetBlockByID(blockID string) (Block, error) {
	_res, _err := _p.client.Call("Explorer.getBlockByID", blockID)
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getBlockByID").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf(Block{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.(Block)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getBlockByID returned invalid type: %v", _t)
			return Block{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return Block{}, _err
}

func (_p ExplorerProxy) GetCoinStatistic() (CoinStatistic, error) {
	_res, _err := _p.client.Call("Explorer.getCoinStatistic")
	if _err == nil {
		_retType := _p.idl.Method("Explorer.getCoinStatistic").Returns
		_res, _err = barrister.Convert(_p.idl, &_retType, reflect.TypeOf(CoinStatistic{}), _res, "")
	}
	if _err == nil {
		_cast, _ok := _res.(CoinStatistic)
		if !_ok {
			_t := reflect.TypeOf(_res)
			_msg := fmt.Sprintf("Explorer.getCoinStatistic returned invalid type: %v", _t)
			return CoinStatistic{}, &barrister.JsonRpcError{Code: -32000, Message: _msg}
		}
		return _cast, nil
	}
	return CoinStatistic{}, _err
}

func NewJSONServer(idl *barrister.Idl, forceASCII bool, explorer Explorer) barrister.Server {
	return NewServer(idl, &barrister.JsonSerializer{forceASCII}, explorer)
}

func NewServer(idl *barrister.Idl, ser barrister.Serializer, explorer Explorer) barrister.Server {
	_svr := barrister.NewServer(idl, ser)
	_svr.AddHandler("Explorer", explorer)
	return _svr
}

var IdlJsonRaw = `[
    {
        "type": "comment",
        "name": "",
        "comment": "",
        "value": "To compile this file:\n1. Install the barrister translator (IDL -\u003e JSON)\nyou need to be root (or use sudo)\npip install barrister",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "comment",
        "name": "",
        "comment": "",
        "value": "2. Install barrister-go\ngo get github.com/coopernurse/barrister-go\ngo install github.com/coopernurse/barrister-go/idl2go",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "comment",
        "name": "",
        "comment": "",
        "value": "3. barrister explorer.idl | $GOPATH/bin/idl2go -i -p explorer",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "CoinStatistic",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "height",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "supply",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "transfers",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "votes",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "aps",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "BlockGenerator",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "name",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "address",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "Block",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "ID",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "height",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "timestamp",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "transfers",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "votes",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "generateBy",
                "type": "BlockGenerator",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "amount",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "forged",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "size",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "Transfer",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "ID",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "sender",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "recipient",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "amount",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "fee",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "timestamp",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "blockID",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "struct",
        "name": "AddressDetails",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": [
            {
                "name": "address",
                "type": "string",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "totalBalance",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            },
            {
                "name": "nonce",
                "type": "int",
                "optional": false,
                "is_array": false,
                "comment": ""
            }
        ],
        "values": null,
        "functions": null,
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "interface",
        "name": "Explorer",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": [
            {
                "name": "getAddressBalance",
                "comment": "get the balance of an address",
                "params": [
                    {
                        "name": "address",
                        "type": "string",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "int",
                    "optional": false,
                    "is_array": false,
                    "comment": ""
                }
            },
            {
                "name": "getAddressDetails",
                "comment": "get the address detail of an iotex address",
                "params": [
                    {
                        "name": "address",
                        "type": "string",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "AddressDetails",
                    "optional": false,
                    "is_array": false,
                    "comment": ""
                }
            },
            {
                "name": "getLastTransfersByRange",
                "comment": "get list of transactions by start block height, transaction offset and limit",
                "params": [
                    {
                        "name": "startBlockHeight",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "offset",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "limit",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "showCoinBase",
                        "type": "bool",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Transfer",
                    "optional": false,
                    "is_array": true,
                    "comment": ""
                }
            },
            {
                "name": "getTransferByID",
                "comment": "get transaction from transaction id",
                "params": [
                    {
                        "name": "transferID",
                        "type": "string",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Transfer",
                    "optional": false,
                    "is_array": false,
                    "comment": ""
                }
            },
            {
                "name": "getTransfersByAddress",
                "comment": "get list of transaction belong to an address",
                "params": [
                    {
                        "name": "address",
                        "type": "string",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Transfer",
                    "optional": false,
                    "is_array": true,
                    "comment": ""
                }
            },
            {
                "name": "getTransfersByBlockID",
                "comment": "get all transfers in a block",
                "params": [
                    {
                        "name": "blockID",
                        "type": "string",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Transfer",
                    "optional": false,
                    "is_array": true,
                    "comment": ""
                }
            },
            {
                "name": "getLastBlocksByRange",
                "comment": "get list of blocks by block id offset and limit",
                "params": [
                    {
                        "name": "offset",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    },
                    {
                        "name": "limit",
                        "type": "int",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Block",
                    "optional": false,
                    "is_array": true,
                    "comment": ""
                }
            },
            {
                "name": "getBlockByID",
                "comment": "get block by block id",
                "params": [
                    {
                        "name": "blockID",
                        "type": "string",
                        "optional": false,
                        "is_array": false,
                        "comment": ""
                    }
                ],
                "returns": {
                    "name": "",
                    "type": "Block",
                    "optional": false,
                    "is_array": false,
                    "comment": ""
                }
            },
            {
                "name": "getCoinStatistic",
                "comment": "get statistic of iotx",
                "params": [],
                "returns": {
                    "name": "",
                    "type": "CoinStatistic",
                    "optional": false,
                    "is_array": false,
                    "comment": ""
                }
            }
        ],
        "barrister_version": "",
        "date_generated": 0,
        "checksum": ""
    },
    {
        "type": "meta",
        "name": "",
        "comment": "",
        "value": "",
        "extends": "",
        "fields": null,
        "values": null,
        "functions": null,
        "barrister_version": "0.1.6",
        "date_generated": 1529472246257,
        "checksum": "5bfc05ce00c6db2acc1370679969d159"
    }
]`