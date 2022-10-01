// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NgmiTreasuryMetaData contains all meta data concerning the NgmiTreasury contract.
var NgmiTreasuryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"payees_\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tier_durations_days_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tier_prices_usdc_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tier_prices_usdt_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tier_prices_busd_\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"master_admin_\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"admin_addresses_\",\"type\":\"address[]\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20PaymentReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriber\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"i_tier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tier_duration_days\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tier_price_in_token\",\"type\":\"uint256\"}],\"name\":\"NewSubscription\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"PayeeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StateReset\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"acceptance\",\"type\":\"bool\"}],\"name\":\"_setConsideredTokenAcceptance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i_tier\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"tier_availability\",\"type\":\"bool\"}],\"name\":\"_setTierAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"}],\"name\":\"acceptTokenForPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateAllTiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i_tier\",\"type\":\"uint256\"}],\"name\":\"activateTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"new_admins\",\"type\":\"address[]\"}],\"name\":\"addAdmins\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_token_address\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"new_token_tier_prices\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"is_new_token_accepted\",\"type\":\"bool\"}],\"name\":\"addConsideredToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"is_new_tier_purchasable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"new_tier_duration_days\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"new_tier_prices\",\"type\":\"uint256[]\"}],\"name\":\"addTier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"admin_addresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"considered_tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractBUSD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractUSDC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractUSDT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateAllTiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i_tier\",\"type\":\"uint256\"}],\"name\":\"deactivateTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"}],\"name\":\"forbidTokenForPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdminAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsideredTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getDaysPassedFromCurrentSubscription\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getDaysRemainingFromCurrentSubscription\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMasterAdminAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfTiers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getSecondsPassedFromCurrentSubscription\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getSecondsRemainingFromCurrentSubscription\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTierAvailabilities\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i_tier\",\"type\":\"uint256\"}],\"name\":\"getTierAvailability\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i_tier\",\"type\":\"uint256\"}],\"name\":\"getTierDurationInDays\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTierDurationsInDays\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i_tier\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"}],\"name\":\"getTierPriceInToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i_tier\",\"type\":\"uint256\"}],\"name\":\"getTierPricesInAllConsideredTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAcceptances\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserFirstSubscriptionTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserLatestSubscriptionTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserTier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i_tier\",\"type\":\"uint256\"}],\"name\":\"isExistingTier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"}],\"name\":\"isTokenAccepted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"}],\"name\":\"isTokenConsidered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isUserSubscribed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"master_admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"payee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"releasable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"releasable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"released\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"released\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"removed_admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"removed_admins\",\"type\":\"address[]\"}],\"name\":\"removeAdmins\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool[]\",\"name\":\"new_is_token_accepted_array\",\"type\":\"bool[]\"}],\"name\":\"setAllConsideredTokensAcceptances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"new_tier_durations_days\",\"type\":\"uint256[]\"}],\"name\":\"setAllTierDurationsInDays\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"first_subscription_timestamps\",\"type\":\"uint256[]\"}],\"name\":\"setMultipleUserFirstSubscriptionTimestamps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"latest_subscription_timestamps\",\"type\":\"uint256[]\"}],\"name\":\"setMultipleUserLatestSubscriptionTimestamps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"new_i_tiers\",\"type\":\"uint256[]\"}],\"name\":\"setMultipleUserTiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"payees\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"shares_\",\"type\":\"uint256[]\"}],\"name\":\"setNewPayees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i_tier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_tier_duration_days\",\"type\":\"uint256\"}],\"name\":\"setTierDurationInDays\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i_tier\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"new_tier_price_token\",\"type\":\"uint256\"}],\"name\":\"setTierPriceInToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i_tier\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"new_tier_price_in_all_tokens\",\"type\":\"uint256[]\"}],\"name\":\"setTierPricesForAllConsideredTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"first_subscription_timestamp\",\"type\":\"uint256\"}],\"name\":\"setUserFirstSubscriptionTimestamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"latest_subscription_timestamp\",\"type\":\"uint256\"}],\"name\":\"setUserLatestSubscriptionTimestamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"new_i_tier\",\"type\":\"uint256\"}],\"name\":\"setUserTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"shares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i_tier\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"}],\"name\":\"subscribe\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tier_availabilities\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tier_durations_days\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tier_prices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"totalReleased\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReleased\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i_new_tier\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"user_first_subscription_timestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"user_latest_subscription_timestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"user_tiers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// NgmiTreasuryABI is the input ABI used to generate the binding from.
// Deprecated: Use NgmiTreasuryMetaData.ABI instead.
var NgmiTreasuryABI = NgmiTreasuryMetaData.ABI

// NgmiTreasury is an auto generated Go binding around an Ethereum contract.
type NgmiTreasury struct {
	NgmiTreasuryCaller     // Read-only binding to the contract
	NgmiTreasuryTransactor // Write-only binding to the contract
	NgmiTreasuryFilterer   // Log filterer for contract events
}

// NgmiTreasuryCaller is an auto generated read-only Go binding around an Ethereum contract.
type NgmiTreasuryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NgmiTreasuryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NgmiTreasuryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NgmiTreasuryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NgmiTreasuryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NgmiTreasurySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NgmiTreasurySession struct {
	Contract     *NgmiTreasury     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NgmiTreasuryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NgmiTreasuryCallerSession struct {
	Contract *NgmiTreasuryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NgmiTreasuryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NgmiTreasuryTransactorSession struct {
	Contract     *NgmiTreasuryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NgmiTreasuryRaw is an auto generated low-level Go binding around an Ethereum contract.
type NgmiTreasuryRaw struct {
	Contract *NgmiTreasury // Generic contract binding to access the raw methods on
}

// NgmiTreasuryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NgmiTreasuryCallerRaw struct {
	Contract *NgmiTreasuryCaller // Generic read-only contract binding to access the raw methods on
}

// NgmiTreasuryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NgmiTreasuryTransactorRaw struct {
	Contract *NgmiTreasuryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNgmiTreasury creates a new instance of NgmiTreasury, bound to a specific deployed contract.
func NewNgmiTreasury(address common.Address, backend bind.ContractBackend) (*NgmiTreasury, error) {
	contract, err := bindNgmiTreasury(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NgmiTreasury{NgmiTreasuryCaller: NgmiTreasuryCaller{contract: contract}, NgmiTreasuryTransactor: NgmiTreasuryTransactor{contract: contract}, NgmiTreasuryFilterer: NgmiTreasuryFilterer{contract: contract}}, nil
}

// NewNgmiTreasuryCaller creates a new read-only instance of NgmiTreasury, bound to a specific deployed contract.
func NewNgmiTreasuryCaller(address common.Address, caller bind.ContractCaller) (*NgmiTreasuryCaller, error) {
	contract, err := bindNgmiTreasury(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NgmiTreasuryCaller{contract: contract}, nil
}

// NewNgmiTreasuryTransactor creates a new write-only instance of NgmiTreasury, bound to a specific deployed contract.
func NewNgmiTreasuryTransactor(address common.Address, transactor bind.ContractTransactor) (*NgmiTreasuryTransactor, error) {
	contract, err := bindNgmiTreasury(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NgmiTreasuryTransactor{contract: contract}, nil
}

// NewNgmiTreasuryFilterer creates a new log filterer instance of NgmiTreasury, bound to a specific deployed contract.
func NewNgmiTreasuryFilterer(address common.Address, filterer bind.ContractFilterer) (*NgmiTreasuryFilterer, error) {
	contract, err := bindNgmiTreasury(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NgmiTreasuryFilterer{contract: contract}, nil
}

// bindNgmiTreasury binds a generic wrapper to an already deployed contract.
func bindNgmiTreasury(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NgmiTreasuryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NgmiTreasury *NgmiTreasuryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NgmiTreasury.Contract.NgmiTreasuryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NgmiTreasury *NgmiTreasuryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.NgmiTreasuryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NgmiTreasury *NgmiTreasuryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.NgmiTreasuryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NgmiTreasury *NgmiTreasuryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NgmiTreasury.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NgmiTreasury *NgmiTreasuryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NgmiTreasury *NgmiTreasuryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.contract.Transact(opts, method, params...)
}

// AdminAddresses is a free data retrieval call binding the contract method 0x4f2b1aca.
//
// Solidity: function admin_addresses(uint256 ) view returns(address)
func (_NgmiTreasury *NgmiTreasuryCaller) AdminAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "admin_addresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminAddresses is a free data retrieval call binding the contract method 0x4f2b1aca.
//
// Solidity: function admin_addresses(uint256 ) view returns(address)
func (_NgmiTreasury *NgmiTreasurySession) AdminAddresses(arg0 *big.Int) (common.Address, error) {
	return _NgmiTreasury.Contract.AdminAddresses(&_NgmiTreasury.CallOpts, arg0)
}

// AdminAddresses is a free data retrieval call binding the contract method 0x4f2b1aca.
//
// Solidity: function admin_addresses(uint256 ) view returns(address)
func (_NgmiTreasury *NgmiTreasuryCallerSession) AdminAddresses(arg0 *big.Int) (common.Address, error) {
	return _NgmiTreasury.Contract.AdminAddresses(&_NgmiTreasury.CallOpts, arg0)
}

// ConsideredTokens is a free data retrieval call binding the contract method 0xf55f731c.
//
// Solidity: function considered_tokens(uint256 ) view returns(address)
func (_NgmiTreasury *NgmiTreasuryCaller) ConsideredTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "considered_tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConsideredTokens is a free data retrieval call binding the contract method 0xf55f731c.
//
// Solidity: function considered_tokens(uint256 ) view returns(address)
func (_NgmiTreasury *NgmiTreasurySession) ConsideredTokens(arg0 *big.Int) (common.Address, error) {
	return _NgmiTreasury.Contract.ConsideredTokens(&_NgmiTreasury.CallOpts, arg0)
}

// ConsideredTokens is a free data retrieval call binding the contract method 0xf55f731c.
//
// Solidity: function considered_tokens(uint256 ) view returns(address)
func (_NgmiTreasury *NgmiTreasuryCallerSession) ConsideredTokens(arg0 *big.Int) (common.Address, error) {
	return _NgmiTreasury.Contract.ConsideredTokens(&_NgmiTreasury.CallOpts, arg0)
}

// ContractBUSD is a free data retrieval call binding the contract method 0xd7700667.
//
// Solidity: function contractBUSD() view returns(address)
func (_NgmiTreasury *NgmiTreasuryCaller) ContractBUSD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "contractBUSD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractBUSD is a free data retrieval call binding the contract method 0xd7700667.
//
// Solidity: function contractBUSD() view returns(address)
func (_NgmiTreasury *NgmiTreasurySession) ContractBUSD() (common.Address, error) {
	return _NgmiTreasury.Contract.ContractBUSD(&_NgmiTreasury.CallOpts)
}

// ContractBUSD is a free data retrieval call binding the contract method 0xd7700667.
//
// Solidity: function contractBUSD() view returns(address)
func (_NgmiTreasury *NgmiTreasuryCallerSession) ContractBUSD() (common.Address, error) {
	return _NgmiTreasury.Contract.ContractBUSD(&_NgmiTreasury.CallOpts)
}

// ContractUSDC is a free data retrieval call binding the contract method 0xc22cbad8.
//
// Solidity: function contractUSDC() view returns(address)
func (_NgmiTreasury *NgmiTreasuryCaller) ContractUSDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "contractUSDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractUSDC is a free data retrieval call binding the contract method 0xc22cbad8.
//
// Solidity: function contractUSDC() view returns(address)
func (_NgmiTreasury *NgmiTreasurySession) ContractUSDC() (common.Address, error) {
	return _NgmiTreasury.Contract.ContractUSDC(&_NgmiTreasury.CallOpts)
}

// ContractUSDC is a free data retrieval call binding the contract method 0xc22cbad8.
//
// Solidity: function contractUSDC() view returns(address)
func (_NgmiTreasury *NgmiTreasuryCallerSession) ContractUSDC() (common.Address, error) {
	return _NgmiTreasury.Contract.ContractUSDC(&_NgmiTreasury.CallOpts)
}

// ContractUSDT is a free data retrieval call binding the contract method 0x4e48c198.
//
// Solidity: function contractUSDT() view returns(address)
func (_NgmiTreasury *NgmiTreasuryCaller) ContractUSDT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "contractUSDT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractUSDT is a free data retrieval call binding the contract method 0x4e48c198.
//
// Solidity: function contractUSDT() view returns(address)
func (_NgmiTreasury *NgmiTreasurySession) ContractUSDT() (common.Address, error) {
	return _NgmiTreasury.Contract.ContractUSDT(&_NgmiTreasury.CallOpts)
}

// ContractUSDT is a free data retrieval call binding the contract method 0x4e48c198.
//
// Solidity: function contractUSDT() view returns(address)
func (_NgmiTreasury *NgmiTreasuryCallerSession) ContractUSDT() (common.Address, error) {
	return _NgmiTreasury.Contract.ContractUSDT(&_NgmiTreasury.CallOpts)
}

// GetAdminAddresses is a free data retrieval call binding the contract method 0x64ddfa29.
//
// Solidity: function getAdminAddresses() view returns(address[])
func (_NgmiTreasury *NgmiTreasuryCaller) GetAdminAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getAdminAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdminAddresses is a free data retrieval call binding the contract method 0x64ddfa29.
//
// Solidity: function getAdminAddresses() view returns(address[])
func (_NgmiTreasury *NgmiTreasurySession) GetAdminAddresses() ([]common.Address, error) {
	return _NgmiTreasury.Contract.GetAdminAddresses(&_NgmiTreasury.CallOpts)
}

// GetAdminAddresses is a free data retrieval call binding the contract method 0x64ddfa29.
//
// Solidity: function getAdminAddresses() view returns(address[])
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetAdminAddresses() ([]common.Address, error) {
	return _NgmiTreasury.Contract.GetAdminAddresses(&_NgmiTreasury.CallOpts)
}

// GetConsideredTokens is a free data retrieval call binding the contract method 0x3bb8a87f.
//
// Solidity: function getConsideredTokens() view returns(address[])
func (_NgmiTreasury *NgmiTreasuryCaller) GetConsideredTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getConsideredTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetConsideredTokens is a free data retrieval call binding the contract method 0x3bb8a87f.
//
// Solidity: function getConsideredTokens() view returns(address[])
func (_NgmiTreasury *NgmiTreasurySession) GetConsideredTokens() ([]common.Address, error) {
	return _NgmiTreasury.Contract.GetConsideredTokens(&_NgmiTreasury.CallOpts)
}

// GetConsideredTokens is a free data retrieval call binding the contract method 0x3bb8a87f.
//
// Solidity: function getConsideredTokens() view returns(address[])
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetConsideredTokens() ([]common.Address, error) {
	return _NgmiTreasury.Contract.GetConsideredTokens(&_NgmiTreasury.CallOpts)
}

// GetDaysPassedFromCurrentSubscription is a free data retrieval call binding the contract method 0x5f9e02a5.
//
// Solidity: function getDaysPassedFromCurrentSubscription(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) GetDaysPassedFromCurrentSubscription(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getDaysPassedFromCurrentSubscription", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDaysPassedFromCurrentSubscription is a free data retrieval call binding the contract method 0x5f9e02a5.
//
// Solidity: function getDaysPassedFromCurrentSubscription(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) GetDaysPassedFromCurrentSubscription(user common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetDaysPassedFromCurrentSubscription(&_NgmiTreasury.CallOpts, user)
}

// GetDaysPassedFromCurrentSubscription is a free data retrieval call binding the contract method 0x5f9e02a5.
//
// Solidity: function getDaysPassedFromCurrentSubscription(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetDaysPassedFromCurrentSubscription(user common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetDaysPassedFromCurrentSubscription(&_NgmiTreasury.CallOpts, user)
}

// GetDaysRemainingFromCurrentSubscription is a free data retrieval call binding the contract method 0x2d4ed3a2.
//
// Solidity: function getDaysRemainingFromCurrentSubscription(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) GetDaysRemainingFromCurrentSubscription(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getDaysRemainingFromCurrentSubscription", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDaysRemainingFromCurrentSubscription is a free data retrieval call binding the contract method 0x2d4ed3a2.
//
// Solidity: function getDaysRemainingFromCurrentSubscription(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) GetDaysRemainingFromCurrentSubscription(user common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetDaysRemainingFromCurrentSubscription(&_NgmiTreasury.CallOpts, user)
}

// GetDaysRemainingFromCurrentSubscription is a free data retrieval call binding the contract method 0x2d4ed3a2.
//
// Solidity: function getDaysRemainingFromCurrentSubscription(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetDaysRemainingFromCurrentSubscription(user common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetDaysRemainingFromCurrentSubscription(&_NgmiTreasury.CallOpts, user)
}

// GetMasterAdminAddress is a free data retrieval call binding the contract method 0x43d5aac7.
//
// Solidity: function getMasterAdminAddress() view returns(address)
func (_NgmiTreasury *NgmiTreasuryCaller) GetMasterAdminAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getMasterAdminAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMasterAdminAddress is a free data retrieval call binding the contract method 0x43d5aac7.
//
// Solidity: function getMasterAdminAddress() view returns(address)
func (_NgmiTreasury *NgmiTreasurySession) GetMasterAdminAddress() (common.Address, error) {
	return _NgmiTreasury.Contract.GetMasterAdminAddress(&_NgmiTreasury.CallOpts)
}

// GetMasterAdminAddress is a free data retrieval call binding the contract method 0x43d5aac7.
//
// Solidity: function getMasterAdminAddress() view returns(address)
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetMasterAdminAddress() (common.Address, error) {
	return _NgmiTreasury.Contract.GetMasterAdminAddress(&_NgmiTreasury.CallOpts)
}

// GetNumberOfTiers is a free data retrieval call binding the contract method 0x3a4fd45c.
//
// Solidity: function getNumberOfTiers() view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) GetNumberOfTiers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getNumberOfTiers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfTiers is a free data retrieval call binding the contract method 0x3a4fd45c.
//
// Solidity: function getNumberOfTiers() view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) GetNumberOfTiers() (*big.Int, error) {
	return _NgmiTreasury.Contract.GetNumberOfTiers(&_NgmiTreasury.CallOpts)
}

// GetNumberOfTiers is a free data retrieval call binding the contract method 0x3a4fd45c.
//
// Solidity: function getNumberOfTiers() view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetNumberOfTiers() (*big.Int, error) {
	return _NgmiTreasury.Contract.GetNumberOfTiers(&_NgmiTreasury.CallOpts)
}

// GetSecondsPassedFromCurrentSubscription is a free data retrieval call binding the contract method 0xb888a1d1.
//
// Solidity: function getSecondsPassedFromCurrentSubscription(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) GetSecondsPassedFromCurrentSubscription(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getSecondsPassedFromCurrentSubscription", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSecondsPassedFromCurrentSubscription is a free data retrieval call binding the contract method 0xb888a1d1.
//
// Solidity: function getSecondsPassedFromCurrentSubscription(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) GetSecondsPassedFromCurrentSubscription(user common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetSecondsPassedFromCurrentSubscription(&_NgmiTreasury.CallOpts, user)
}

// GetSecondsPassedFromCurrentSubscription is a free data retrieval call binding the contract method 0xb888a1d1.
//
// Solidity: function getSecondsPassedFromCurrentSubscription(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetSecondsPassedFromCurrentSubscription(user common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetSecondsPassedFromCurrentSubscription(&_NgmiTreasury.CallOpts, user)
}

// GetSecondsRemainingFromCurrentSubscription is a free data retrieval call binding the contract method 0xb5dbc0cc.
//
// Solidity: function getSecondsRemainingFromCurrentSubscription(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) GetSecondsRemainingFromCurrentSubscription(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getSecondsRemainingFromCurrentSubscription", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSecondsRemainingFromCurrentSubscription is a free data retrieval call binding the contract method 0xb5dbc0cc.
//
// Solidity: function getSecondsRemainingFromCurrentSubscription(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) GetSecondsRemainingFromCurrentSubscription(user common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetSecondsRemainingFromCurrentSubscription(&_NgmiTreasury.CallOpts, user)
}

// GetSecondsRemainingFromCurrentSubscription is a free data retrieval call binding the contract method 0xb5dbc0cc.
//
// Solidity: function getSecondsRemainingFromCurrentSubscription(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetSecondsRemainingFromCurrentSubscription(user common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetSecondsRemainingFromCurrentSubscription(&_NgmiTreasury.CallOpts, user)
}

// GetTierAvailabilities is a free data retrieval call binding the contract method 0x955d0f2c.
//
// Solidity: function getTierAvailabilities() view returns(bool[])
func (_NgmiTreasury *NgmiTreasuryCaller) GetTierAvailabilities(opts *bind.CallOpts) ([]bool, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getTierAvailabilities")

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// GetTierAvailabilities is a free data retrieval call binding the contract method 0x955d0f2c.
//
// Solidity: function getTierAvailabilities() view returns(bool[])
func (_NgmiTreasury *NgmiTreasurySession) GetTierAvailabilities() ([]bool, error) {
	return _NgmiTreasury.Contract.GetTierAvailabilities(&_NgmiTreasury.CallOpts)
}

// GetTierAvailabilities is a free data retrieval call binding the contract method 0x955d0f2c.
//
// Solidity: function getTierAvailabilities() view returns(bool[])
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetTierAvailabilities() ([]bool, error) {
	return _NgmiTreasury.Contract.GetTierAvailabilities(&_NgmiTreasury.CallOpts)
}

// GetTierAvailability is a free data retrieval call binding the contract method 0xdaf54969.
//
// Solidity: function getTierAvailability(uint256 i_tier) view returns(bool)
func (_NgmiTreasury *NgmiTreasuryCaller) GetTierAvailability(opts *bind.CallOpts, i_tier *big.Int) (bool, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getTierAvailability", i_tier)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetTierAvailability is a free data retrieval call binding the contract method 0xdaf54969.
//
// Solidity: function getTierAvailability(uint256 i_tier) view returns(bool)
func (_NgmiTreasury *NgmiTreasurySession) GetTierAvailability(i_tier *big.Int) (bool, error) {
	return _NgmiTreasury.Contract.GetTierAvailability(&_NgmiTreasury.CallOpts, i_tier)
}

// GetTierAvailability is a free data retrieval call binding the contract method 0xdaf54969.
//
// Solidity: function getTierAvailability(uint256 i_tier) view returns(bool)
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetTierAvailability(i_tier *big.Int) (bool, error) {
	return _NgmiTreasury.Contract.GetTierAvailability(&_NgmiTreasury.CallOpts, i_tier)
}

// GetTierDurationInDays is a free data retrieval call binding the contract method 0x0aa5229d.
//
// Solidity: function getTierDurationInDays(uint256 i_tier) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) GetTierDurationInDays(opts *bind.CallOpts, i_tier *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getTierDurationInDays", i_tier)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTierDurationInDays is a free data retrieval call binding the contract method 0x0aa5229d.
//
// Solidity: function getTierDurationInDays(uint256 i_tier) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) GetTierDurationInDays(i_tier *big.Int) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetTierDurationInDays(&_NgmiTreasury.CallOpts, i_tier)
}

// GetTierDurationInDays is a free data retrieval call binding the contract method 0x0aa5229d.
//
// Solidity: function getTierDurationInDays(uint256 i_tier) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetTierDurationInDays(i_tier *big.Int) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetTierDurationInDays(&_NgmiTreasury.CallOpts, i_tier)
}

// GetTierDurationsInDays is a free data retrieval call binding the contract method 0x75a4d678.
//
// Solidity: function getTierDurationsInDays() view returns(uint256[])
func (_NgmiTreasury *NgmiTreasuryCaller) GetTierDurationsInDays(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getTierDurationsInDays")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTierDurationsInDays is a free data retrieval call binding the contract method 0x75a4d678.
//
// Solidity: function getTierDurationsInDays() view returns(uint256[])
func (_NgmiTreasury *NgmiTreasurySession) GetTierDurationsInDays() ([]*big.Int, error) {
	return _NgmiTreasury.Contract.GetTierDurationsInDays(&_NgmiTreasury.CallOpts)
}

// GetTierDurationsInDays is a free data retrieval call binding the contract method 0x75a4d678.
//
// Solidity: function getTierDurationsInDays() view returns(uint256[])
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetTierDurationsInDays() ([]*big.Int, error) {
	return _NgmiTreasury.Contract.GetTierDurationsInDays(&_NgmiTreasury.CallOpts)
}

// GetTierPriceInToken is a free data retrieval call binding the contract method 0x94d84da2.
//
// Solidity: function getTierPriceInToken(uint256 i_tier, address token_address) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) GetTierPriceInToken(opts *bind.CallOpts, i_tier *big.Int, token_address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getTierPriceInToken", i_tier, token_address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTierPriceInToken is a free data retrieval call binding the contract method 0x94d84da2.
//
// Solidity: function getTierPriceInToken(uint256 i_tier, address token_address) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) GetTierPriceInToken(i_tier *big.Int, token_address common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetTierPriceInToken(&_NgmiTreasury.CallOpts, i_tier, token_address)
}

// GetTierPriceInToken is a free data retrieval call binding the contract method 0x94d84da2.
//
// Solidity: function getTierPriceInToken(uint256 i_tier, address token_address) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetTierPriceInToken(i_tier *big.Int, token_address common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetTierPriceInToken(&_NgmiTreasury.CallOpts, i_tier, token_address)
}

// GetTierPricesInAllConsideredTokens is a free data retrieval call binding the contract method 0x470aa21b.
//
// Solidity: function getTierPricesInAllConsideredTokens(uint256 i_tier) view returns(uint256[])
func (_NgmiTreasury *NgmiTreasuryCaller) GetTierPricesInAllConsideredTokens(opts *bind.CallOpts, i_tier *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getTierPricesInAllConsideredTokens", i_tier)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTierPricesInAllConsideredTokens is a free data retrieval call binding the contract method 0x470aa21b.
//
// Solidity: function getTierPricesInAllConsideredTokens(uint256 i_tier) view returns(uint256[])
func (_NgmiTreasury *NgmiTreasurySession) GetTierPricesInAllConsideredTokens(i_tier *big.Int) ([]*big.Int, error) {
	return _NgmiTreasury.Contract.GetTierPricesInAllConsideredTokens(&_NgmiTreasury.CallOpts, i_tier)
}

// GetTierPricesInAllConsideredTokens is a free data retrieval call binding the contract method 0x470aa21b.
//
// Solidity: function getTierPricesInAllConsideredTokens(uint256 i_tier) view returns(uint256[])
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetTierPricesInAllConsideredTokens(i_tier *big.Int) ([]*big.Int, error) {
	return _NgmiTreasury.Contract.GetTierPricesInAllConsideredTokens(&_NgmiTreasury.CallOpts, i_tier)
}

// GetTokenAcceptances is a free data retrieval call binding the contract method 0xf12fb085.
//
// Solidity: function getTokenAcceptances() view returns(bool[])
func (_NgmiTreasury *NgmiTreasuryCaller) GetTokenAcceptances(opts *bind.CallOpts) ([]bool, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getTokenAcceptances")

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// GetTokenAcceptances is a free data retrieval call binding the contract method 0xf12fb085.
//
// Solidity: function getTokenAcceptances() view returns(bool[])
func (_NgmiTreasury *NgmiTreasurySession) GetTokenAcceptances() ([]bool, error) {
	return _NgmiTreasury.Contract.GetTokenAcceptances(&_NgmiTreasury.CallOpts)
}

// GetTokenAcceptances is a free data retrieval call binding the contract method 0xf12fb085.
//
// Solidity: function getTokenAcceptances() view returns(bool[])
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetTokenAcceptances() ([]bool, error) {
	return _NgmiTreasury.Contract.GetTokenAcceptances(&_NgmiTreasury.CallOpts)
}

// GetUserFirstSubscriptionTimestamp is a free data retrieval call binding the contract method 0xad74b699.
//
// Solidity: function getUserFirstSubscriptionTimestamp(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) GetUserFirstSubscriptionTimestamp(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getUserFirstSubscriptionTimestamp", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserFirstSubscriptionTimestamp is a free data retrieval call binding the contract method 0xad74b699.
//
// Solidity: function getUserFirstSubscriptionTimestamp(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) GetUserFirstSubscriptionTimestamp(user common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetUserFirstSubscriptionTimestamp(&_NgmiTreasury.CallOpts, user)
}

// GetUserFirstSubscriptionTimestamp is a free data retrieval call binding the contract method 0xad74b699.
//
// Solidity: function getUserFirstSubscriptionTimestamp(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetUserFirstSubscriptionTimestamp(user common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetUserFirstSubscriptionTimestamp(&_NgmiTreasury.CallOpts, user)
}

// GetUserLatestSubscriptionTimestamp is a free data retrieval call binding the contract method 0xe172e8f9.
//
// Solidity: function getUserLatestSubscriptionTimestamp(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) GetUserLatestSubscriptionTimestamp(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getUserLatestSubscriptionTimestamp", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLatestSubscriptionTimestamp is a free data retrieval call binding the contract method 0xe172e8f9.
//
// Solidity: function getUserLatestSubscriptionTimestamp(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) GetUserLatestSubscriptionTimestamp(user common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetUserLatestSubscriptionTimestamp(&_NgmiTreasury.CallOpts, user)
}

// GetUserLatestSubscriptionTimestamp is a free data retrieval call binding the contract method 0xe172e8f9.
//
// Solidity: function getUserLatestSubscriptionTimestamp(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetUserLatestSubscriptionTimestamp(user common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetUserLatestSubscriptionTimestamp(&_NgmiTreasury.CallOpts, user)
}

// GetUserTier is a free data retrieval call binding the contract method 0xe4d2620e.
//
// Solidity: function getUserTier(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) GetUserTier(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "getUserTier", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserTier is a free data retrieval call binding the contract method 0xe4d2620e.
//
// Solidity: function getUserTier(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) GetUserTier(user common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetUserTier(&_NgmiTreasury.CallOpts, user)
}

// GetUserTier is a free data retrieval call binding the contract method 0xe4d2620e.
//
// Solidity: function getUserTier(address user) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) GetUserTier(user common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.GetUserTier(&_NgmiTreasury.CallOpts, user)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address user) view returns(bool)
func (_NgmiTreasury *NgmiTreasuryCaller) IsAdmin(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "isAdmin", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address user) view returns(bool)
func (_NgmiTreasury *NgmiTreasurySession) IsAdmin(user common.Address) (bool, error) {
	return _NgmiTreasury.Contract.IsAdmin(&_NgmiTreasury.CallOpts, user)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address user) view returns(bool)
func (_NgmiTreasury *NgmiTreasuryCallerSession) IsAdmin(user common.Address) (bool, error) {
	return _NgmiTreasury.Contract.IsAdmin(&_NgmiTreasury.CallOpts, user)
}

// IsExistingTier is a free data retrieval call binding the contract method 0xa63f3c8c.
//
// Solidity: function isExistingTier(uint256 i_tier) view returns(bool)
func (_NgmiTreasury *NgmiTreasuryCaller) IsExistingTier(opts *bind.CallOpts, i_tier *big.Int) (bool, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "isExistingTier", i_tier)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistingTier is a free data retrieval call binding the contract method 0xa63f3c8c.
//
// Solidity: function isExistingTier(uint256 i_tier) view returns(bool)
func (_NgmiTreasury *NgmiTreasurySession) IsExistingTier(i_tier *big.Int) (bool, error) {
	return _NgmiTreasury.Contract.IsExistingTier(&_NgmiTreasury.CallOpts, i_tier)
}

// IsExistingTier is a free data retrieval call binding the contract method 0xa63f3c8c.
//
// Solidity: function isExistingTier(uint256 i_tier) view returns(bool)
func (_NgmiTreasury *NgmiTreasuryCallerSession) IsExistingTier(i_tier *big.Int) (bool, error) {
	return _NgmiTreasury.Contract.IsExistingTier(&_NgmiTreasury.CallOpts, i_tier)
}

// IsTokenAccepted is a free data retrieval call binding the contract method 0x0560bd96.
//
// Solidity: function isTokenAccepted(address token_address) view returns(bool)
func (_NgmiTreasury *NgmiTreasuryCaller) IsTokenAccepted(opts *bind.CallOpts, token_address common.Address) (bool, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "isTokenAccepted", token_address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenAccepted is a free data retrieval call binding the contract method 0x0560bd96.
//
// Solidity: function isTokenAccepted(address token_address) view returns(bool)
func (_NgmiTreasury *NgmiTreasurySession) IsTokenAccepted(token_address common.Address) (bool, error) {
	return _NgmiTreasury.Contract.IsTokenAccepted(&_NgmiTreasury.CallOpts, token_address)
}

// IsTokenAccepted is a free data retrieval call binding the contract method 0x0560bd96.
//
// Solidity: function isTokenAccepted(address token_address) view returns(bool)
func (_NgmiTreasury *NgmiTreasuryCallerSession) IsTokenAccepted(token_address common.Address) (bool, error) {
	return _NgmiTreasury.Contract.IsTokenAccepted(&_NgmiTreasury.CallOpts, token_address)
}

// IsTokenConsidered is a free data retrieval call binding the contract method 0xc1ed476e.
//
// Solidity: function isTokenConsidered(address token_address) view returns(bool)
func (_NgmiTreasury *NgmiTreasuryCaller) IsTokenConsidered(opts *bind.CallOpts, token_address common.Address) (bool, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "isTokenConsidered", token_address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenConsidered is a free data retrieval call binding the contract method 0xc1ed476e.
//
// Solidity: function isTokenConsidered(address token_address) view returns(bool)
func (_NgmiTreasury *NgmiTreasurySession) IsTokenConsidered(token_address common.Address) (bool, error) {
	return _NgmiTreasury.Contract.IsTokenConsidered(&_NgmiTreasury.CallOpts, token_address)
}

// IsTokenConsidered is a free data retrieval call binding the contract method 0xc1ed476e.
//
// Solidity: function isTokenConsidered(address token_address) view returns(bool)
func (_NgmiTreasury *NgmiTreasuryCallerSession) IsTokenConsidered(token_address common.Address) (bool, error) {
	return _NgmiTreasury.Contract.IsTokenConsidered(&_NgmiTreasury.CallOpts, token_address)
}

// IsUserSubscribed is a free data retrieval call binding the contract method 0xabf0eb50.
//
// Solidity: function isUserSubscribed(address user) view returns(bool)
func (_NgmiTreasury *NgmiTreasuryCaller) IsUserSubscribed(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "isUserSubscribed", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserSubscribed is a free data retrieval call binding the contract method 0xabf0eb50.
//
// Solidity: function isUserSubscribed(address user) view returns(bool)
func (_NgmiTreasury *NgmiTreasurySession) IsUserSubscribed(user common.Address) (bool, error) {
	return _NgmiTreasury.Contract.IsUserSubscribed(&_NgmiTreasury.CallOpts, user)
}

// IsUserSubscribed is a free data retrieval call binding the contract method 0xabf0eb50.
//
// Solidity: function isUserSubscribed(address user) view returns(bool)
func (_NgmiTreasury *NgmiTreasuryCallerSession) IsUserSubscribed(user common.Address) (bool, error) {
	return _NgmiTreasury.Contract.IsUserSubscribed(&_NgmiTreasury.CallOpts, user)
}

// MasterAdmin is a free data retrieval call binding the contract method 0x83fe9ab0.
//
// Solidity: function master_admin() view returns(address)
func (_NgmiTreasury *NgmiTreasuryCaller) MasterAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "master_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterAdmin is a free data retrieval call binding the contract method 0x83fe9ab0.
//
// Solidity: function master_admin() view returns(address)
func (_NgmiTreasury *NgmiTreasurySession) MasterAdmin() (common.Address, error) {
	return _NgmiTreasury.Contract.MasterAdmin(&_NgmiTreasury.CallOpts)
}

// MasterAdmin is a free data retrieval call binding the contract method 0x83fe9ab0.
//
// Solidity: function master_admin() view returns(address)
func (_NgmiTreasury *NgmiTreasuryCallerSession) MasterAdmin() (common.Address, error) {
	return _NgmiTreasury.Contract.MasterAdmin(&_NgmiTreasury.CallOpts)
}

// Payee is a free data retrieval call binding the contract method 0x8b83209b.
//
// Solidity: function payee(uint256 index) view returns(address)
func (_NgmiTreasury *NgmiTreasuryCaller) Payee(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "payee", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Payee is a free data retrieval call binding the contract method 0x8b83209b.
//
// Solidity: function payee(uint256 index) view returns(address)
func (_NgmiTreasury *NgmiTreasurySession) Payee(index *big.Int) (common.Address, error) {
	return _NgmiTreasury.Contract.Payee(&_NgmiTreasury.CallOpts, index)
}

// Payee is a free data retrieval call binding the contract method 0x8b83209b.
//
// Solidity: function payee(uint256 index) view returns(address)
func (_NgmiTreasury *NgmiTreasuryCallerSession) Payee(index *big.Int) (common.Address, error) {
	return _NgmiTreasury.Contract.Payee(&_NgmiTreasury.CallOpts, index)
}

// Releasable is a free data retrieval call binding the contract method 0xa3f8eace.
//
// Solidity: function releasable(address account) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) Releasable(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "releasable", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Releasable is a free data retrieval call binding the contract method 0xa3f8eace.
//
// Solidity: function releasable(address account) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) Releasable(account common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.Releasable(&_NgmiTreasury.CallOpts, account)
}

// Releasable is a free data retrieval call binding the contract method 0xa3f8eace.
//
// Solidity: function releasable(address account) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) Releasable(account common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.Releasable(&_NgmiTreasury.CallOpts, account)
}

// Releasable0 is a free data retrieval call binding the contract method 0xc45ac050.
//
// Solidity: function releasable(address token, address account) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) Releasable0(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "releasable0", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Releasable0 is a free data retrieval call binding the contract method 0xc45ac050.
//
// Solidity: function releasable(address token, address account) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) Releasable0(token common.Address, account common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.Releasable0(&_NgmiTreasury.CallOpts, token, account)
}

// Releasable0 is a free data retrieval call binding the contract method 0xc45ac050.
//
// Solidity: function releasable(address token, address account) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) Releasable0(token common.Address, account common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.Releasable0(&_NgmiTreasury.CallOpts, token, account)
}

// Released is a free data retrieval call binding the contract method 0x406072a9.
//
// Solidity: function released(address token, address account) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) Released(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "released", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Released is a free data retrieval call binding the contract method 0x406072a9.
//
// Solidity: function released(address token, address account) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) Released(token common.Address, account common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.Released(&_NgmiTreasury.CallOpts, token, account)
}

// Released is a free data retrieval call binding the contract method 0x406072a9.
//
// Solidity: function released(address token, address account) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) Released(token common.Address, account common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.Released(&_NgmiTreasury.CallOpts, token, account)
}

// Released0 is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address account) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) Released0(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "released0", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Released0 is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address account) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) Released0(account common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.Released0(&_NgmiTreasury.CallOpts, account)
}

// Released0 is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address account) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) Released0(account common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.Released0(&_NgmiTreasury.CallOpts, account)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address account) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) Shares(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "shares", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address account) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) Shares(account common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.Shares(&_NgmiTreasury.CallOpts, account)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address account) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) Shares(account common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.Shares(&_NgmiTreasury.CallOpts, account)
}

// TierAvailabilities is a free data retrieval call binding the contract method 0x2d8905c3.
//
// Solidity: function tier_availabilities(uint256 ) view returns(bool)
func (_NgmiTreasury *NgmiTreasuryCaller) TierAvailabilities(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "tier_availabilities", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TierAvailabilities is a free data retrieval call binding the contract method 0x2d8905c3.
//
// Solidity: function tier_availabilities(uint256 ) view returns(bool)
func (_NgmiTreasury *NgmiTreasurySession) TierAvailabilities(arg0 *big.Int) (bool, error) {
	return _NgmiTreasury.Contract.TierAvailabilities(&_NgmiTreasury.CallOpts, arg0)
}

// TierAvailabilities is a free data retrieval call binding the contract method 0x2d8905c3.
//
// Solidity: function tier_availabilities(uint256 ) view returns(bool)
func (_NgmiTreasury *NgmiTreasuryCallerSession) TierAvailabilities(arg0 *big.Int) (bool, error) {
	return _NgmiTreasury.Contract.TierAvailabilities(&_NgmiTreasury.CallOpts, arg0)
}

// TierDurationsDays is a free data retrieval call binding the contract method 0xc02eb210.
//
// Solidity: function tier_durations_days(uint256 ) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) TierDurationsDays(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "tier_durations_days", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierDurationsDays is a free data retrieval call binding the contract method 0xc02eb210.
//
// Solidity: function tier_durations_days(uint256 ) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) TierDurationsDays(arg0 *big.Int) (*big.Int, error) {
	return _NgmiTreasury.Contract.TierDurationsDays(&_NgmiTreasury.CallOpts, arg0)
}

// TierDurationsDays is a free data retrieval call binding the contract method 0xc02eb210.
//
// Solidity: function tier_durations_days(uint256 ) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) TierDurationsDays(arg0 *big.Int) (*big.Int, error) {
	return _NgmiTreasury.Contract.TierDurationsDays(&_NgmiTreasury.CallOpts, arg0)
}

// TierPrices is a free data retrieval call binding the contract method 0x19bcde15.
//
// Solidity: function tier_prices(uint256 , address ) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) TierPrices(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "tier_prices", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TierPrices is a free data retrieval call binding the contract method 0x19bcde15.
//
// Solidity: function tier_prices(uint256 , address ) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) TierPrices(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.TierPrices(&_NgmiTreasury.CallOpts, arg0, arg1)
}

// TierPrices is a free data retrieval call binding the contract method 0x19bcde15.
//
// Solidity: function tier_prices(uint256 , address ) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) TierPrices(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.TierPrices(&_NgmiTreasury.CallOpts, arg0, arg1)
}

// TotalReleased is a free data retrieval call binding the contract method 0xd79779b2.
//
// Solidity: function totalReleased(address token) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) TotalReleased(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "totalReleased", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReleased is a free data retrieval call binding the contract method 0xd79779b2.
//
// Solidity: function totalReleased(address token) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) TotalReleased(token common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.TotalReleased(&_NgmiTreasury.CallOpts, token)
}

// TotalReleased is a free data retrieval call binding the contract method 0xd79779b2.
//
// Solidity: function totalReleased(address token) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) TotalReleased(token common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.TotalReleased(&_NgmiTreasury.CallOpts, token)
}

// TotalReleased0 is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) TotalReleased0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "totalReleased0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReleased0 is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) TotalReleased0() (*big.Int, error) {
	return _NgmiTreasury.Contract.TotalReleased0(&_NgmiTreasury.CallOpts)
}

// TotalReleased0 is a free data retrieval call binding the contract method 0xe33b7de3.
//
// Solidity: function totalReleased() view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) TotalReleased0() (*big.Int, error) {
	return _NgmiTreasury.Contract.TotalReleased0(&_NgmiTreasury.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) TotalShares() (*big.Int, error) {
	return _NgmiTreasury.Contract.TotalShares(&_NgmiTreasury.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) TotalShares() (*big.Int, error) {
	return _NgmiTreasury.Contract.TotalShares(&_NgmiTreasury.CallOpts)
}

// UserFirstSubscriptionTimestamps is a free data retrieval call binding the contract method 0x9cb599fd.
//
// Solidity: function user_first_subscription_timestamps(address ) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) UserFirstSubscriptionTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "user_first_subscription_timestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserFirstSubscriptionTimestamps is a free data retrieval call binding the contract method 0x9cb599fd.
//
// Solidity: function user_first_subscription_timestamps(address ) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) UserFirstSubscriptionTimestamps(arg0 common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.UserFirstSubscriptionTimestamps(&_NgmiTreasury.CallOpts, arg0)
}

// UserFirstSubscriptionTimestamps is a free data retrieval call binding the contract method 0x9cb599fd.
//
// Solidity: function user_first_subscription_timestamps(address ) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) UserFirstSubscriptionTimestamps(arg0 common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.UserFirstSubscriptionTimestamps(&_NgmiTreasury.CallOpts, arg0)
}

// UserLatestSubscriptionTimestamps is a free data retrieval call binding the contract method 0x6c07e5eb.
//
// Solidity: function user_latest_subscription_timestamps(address ) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) UserLatestSubscriptionTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "user_latest_subscription_timestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLatestSubscriptionTimestamps is a free data retrieval call binding the contract method 0x6c07e5eb.
//
// Solidity: function user_latest_subscription_timestamps(address ) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) UserLatestSubscriptionTimestamps(arg0 common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.UserLatestSubscriptionTimestamps(&_NgmiTreasury.CallOpts, arg0)
}

// UserLatestSubscriptionTimestamps is a free data retrieval call binding the contract method 0x6c07e5eb.
//
// Solidity: function user_latest_subscription_timestamps(address ) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) UserLatestSubscriptionTimestamps(arg0 common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.UserLatestSubscriptionTimestamps(&_NgmiTreasury.CallOpts, arg0)
}

// UserTiers is a free data retrieval call binding the contract method 0x2b6d31f5.
//
// Solidity: function user_tiers(address ) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCaller) UserTiers(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NgmiTreasury.contract.Call(opts, &out, "user_tiers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserTiers is a free data retrieval call binding the contract method 0x2b6d31f5.
//
// Solidity: function user_tiers(address ) view returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) UserTiers(arg0 common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.UserTiers(&_NgmiTreasury.CallOpts, arg0)
}

// UserTiers is a free data retrieval call binding the contract method 0x2b6d31f5.
//
// Solidity: function user_tiers(address ) view returns(uint256)
func (_NgmiTreasury *NgmiTreasuryCallerSession) UserTiers(arg0 common.Address) (*big.Int, error) {
	return _NgmiTreasury.Contract.UserTiers(&_NgmiTreasury.CallOpts, arg0)
}

// SetConsideredTokenAcceptance is a paid mutator transaction binding the contract method 0x44ffe5d2.
//
// Solidity: function _setConsideredTokenAcceptance(address token_address, bool acceptance) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) SetConsideredTokenAcceptance(opts *bind.TransactOpts, token_address common.Address, acceptance bool) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "_setConsideredTokenAcceptance", token_address, acceptance)
}

// SetConsideredTokenAcceptance is a paid mutator transaction binding the contract method 0x44ffe5d2.
//
// Solidity: function _setConsideredTokenAcceptance(address token_address, bool acceptance) returns()
func (_NgmiTreasury *NgmiTreasurySession) SetConsideredTokenAcceptance(token_address common.Address, acceptance bool) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetConsideredTokenAcceptance(&_NgmiTreasury.TransactOpts, token_address, acceptance)
}

// SetConsideredTokenAcceptance is a paid mutator transaction binding the contract method 0x44ffe5d2.
//
// Solidity: function _setConsideredTokenAcceptance(address token_address, bool acceptance) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) SetConsideredTokenAcceptance(token_address common.Address, acceptance bool) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetConsideredTokenAcceptance(&_NgmiTreasury.TransactOpts, token_address, acceptance)
}

// SetTierAvailability is a paid mutator transaction binding the contract method 0x6d6d5bb2.
//
// Solidity: function _setTierAvailability(uint256 i_tier, bool tier_availability) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) SetTierAvailability(opts *bind.TransactOpts, i_tier *big.Int, tier_availability bool) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "_setTierAvailability", i_tier, tier_availability)
}

// SetTierAvailability is a paid mutator transaction binding the contract method 0x6d6d5bb2.
//
// Solidity: function _setTierAvailability(uint256 i_tier, bool tier_availability) returns()
func (_NgmiTreasury *NgmiTreasurySession) SetTierAvailability(i_tier *big.Int, tier_availability bool) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetTierAvailability(&_NgmiTreasury.TransactOpts, i_tier, tier_availability)
}

// SetTierAvailability is a paid mutator transaction binding the contract method 0x6d6d5bb2.
//
// Solidity: function _setTierAvailability(uint256 i_tier, bool tier_availability) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) SetTierAvailability(i_tier *big.Int, tier_availability bool) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetTierAvailability(&_NgmiTreasury.TransactOpts, i_tier, tier_availability)
}

// AcceptTokenForPayment is a paid mutator transaction binding the contract method 0x11c9605b.
//
// Solidity: function acceptTokenForPayment(address token_address) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) AcceptTokenForPayment(opts *bind.TransactOpts, token_address common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "acceptTokenForPayment", token_address)
}

// AcceptTokenForPayment is a paid mutator transaction binding the contract method 0x11c9605b.
//
// Solidity: function acceptTokenForPayment(address token_address) returns()
func (_NgmiTreasury *NgmiTreasurySession) AcceptTokenForPayment(token_address common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.AcceptTokenForPayment(&_NgmiTreasury.TransactOpts, token_address)
}

// AcceptTokenForPayment is a paid mutator transaction binding the contract method 0x11c9605b.
//
// Solidity: function acceptTokenForPayment(address token_address) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) AcceptTokenForPayment(token_address common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.AcceptTokenForPayment(&_NgmiTreasury.TransactOpts, token_address)
}

// ActivateAllTiers is a paid mutator transaction binding the contract method 0x674b0718.
//
// Solidity: function activateAllTiers() returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) ActivateAllTiers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "activateAllTiers")
}

// ActivateAllTiers is a paid mutator transaction binding the contract method 0x674b0718.
//
// Solidity: function activateAllTiers() returns()
func (_NgmiTreasury *NgmiTreasurySession) ActivateAllTiers() (*types.Transaction, error) {
	return _NgmiTreasury.Contract.ActivateAllTiers(&_NgmiTreasury.TransactOpts)
}

// ActivateAllTiers is a paid mutator transaction binding the contract method 0x674b0718.
//
// Solidity: function activateAllTiers() returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) ActivateAllTiers() (*types.Transaction, error) {
	return _NgmiTreasury.Contract.ActivateAllTiers(&_NgmiTreasury.TransactOpts)
}

// ActivateTier is a paid mutator transaction binding the contract method 0xdf62c19a.
//
// Solidity: function activateTier(uint256 i_tier) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) ActivateTier(opts *bind.TransactOpts, i_tier *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "activateTier", i_tier)
}

// ActivateTier is a paid mutator transaction binding the contract method 0xdf62c19a.
//
// Solidity: function activateTier(uint256 i_tier) returns()
func (_NgmiTreasury *NgmiTreasurySession) ActivateTier(i_tier *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.ActivateTier(&_NgmiTreasury.TransactOpts, i_tier)
}

// ActivateTier is a paid mutator transaction binding the contract method 0xdf62c19a.
//
// Solidity: function activateTier(uint256 i_tier) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) ActivateTier(i_tier *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.ActivateTier(&_NgmiTreasury.TransactOpts, i_tier)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address new_admin) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) AddAdmin(opts *bind.TransactOpts, new_admin common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "addAdmin", new_admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address new_admin) returns()
func (_NgmiTreasury *NgmiTreasurySession) AddAdmin(new_admin common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.AddAdmin(&_NgmiTreasury.TransactOpts, new_admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address new_admin) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) AddAdmin(new_admin common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.AddAdmin(&_NgmiTreasury.TransactOpts, new_admin)
}

// AddAdmins is a paid mutator transaction binding the contract method 0x9c54df64.
//
// Solidity: function addAdmins(address[] new_admins) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) AddAdmins(opts *bind.TransactOpts, new_admins []common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "addAdmins", new_admins)
}

// AddAdmins is a paid mutator transaction binding the contract method 0x9c54df64.
//
// Solidity: function addAdmins(address[] new_admins) returns()
func (_NgmiTreasury *NgmiTreasurySession) AddAdmins(new_admins []common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.AddAdmins(&_NgmiTreasury.TransactOpts, new_admins)
}

// AddAdmins is a paid mutator transaction binding the contract method 0x9c54df64.
//
// Solidity: function addAdmins(address[] new_admins) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) AddAdmins(new_admins []common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.AddAdmins(&_NgmiTreasury.TransactOpts, new_admins)
}

// AddConsideredToken is a paid mutator transaction binding the contract method 0x5f03ce19.
//
// Solidity: function addConsideredToken(address new_token_address, uint256[] new_token_tier_prices, bool is_new_token_accepted) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) AddConsideredToken(opts *bind.TransactOpts, new_token_address common.Address, new_token_tier_prices []*big.Int, is_new_token_accepted bool) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "addConsideredToken", new_token_address, new_token_tier_prices, is_new_token_accepted)
}

// AddConsideredToken is a paid mutator transaction binding the contract method 0x5f03ce19.
//
// Solidity: function addConsideredToken(address new_token_address, uint256[] new_token_tier_prices, bool is_new_token_accepted) returns()
func (_NgmiTreasury *NgmiTreasurySession) AddConsideredToken(new_token_address common.Address, new_token_tier_prices []*big.Int, is_new_token_accepted bool) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.AddConsideredToken(&_NgmiTreasury.TransactOpts, new_token_address, new_token_tier_prices, is_new_token_accepted)
}

// AddConsideredToken is a paid mutator transaction binding the contract method 0x5f03ce19.
//
// Solidity: function addConsideredToken(address new_token_address, uint256[] new_token_tier_prices, bool is_new_token_accepted) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) AddConsideredToken(new_token_address common.Address, new_token_tier_prices []*big.Int, is_new_token_accepted bool) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.AddConsideredToken(&_NgmiTreasury.TransactOpts, new_token_address, new_token_tier_prices, is_new_token_accepted)
}

// AddTier is a paid mutator transaction binding the contract method 0xc4cdc446.
//
// Solidity: function addTier(bool is_new_tier_purchasable, uint256 new_tier_duration_days, uint256[] new_tier_prices) returns(uint256)
func (_NgmiTreasury *NgmiTreasuryTransactor) AddTier(opts *bind.TransactOpts, is_new_tier_purchasable bool, new_tier_duration_days *big.Int, new_tier_prices []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "addTier", is_new_tier_purchasable, new_tier_duration_days, new_tier_prices)
}

// AddTier is a paid mutator transaction binding the contract method 0xc4cdc446.
//
// Solidity: function addTier(bool is_new_tier_purchasable, uint256 new_tier_duration_days, uint256[] new_tier_prices) returns(uint256)
func (_NgmiTreasury *NgmiTreasurySession) AddTier(is_new_tier_purchasable bool, new_tier_duration_days *big.Int, new_tier_prices []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.AddTier(&_NgmiTreasury.TransactOpts, is_new_tier_purchasable, new_tier_duration_days, new_tier_prices)
}

// AddTier is a paid mutator transaction binding the contract method 0xc4cdc446.
//
// Solidity: function addTier(bool is_new_tier_purchasable, uint256 new_tier_duration_days, uint256[] new_tier_prices) returns(uint256)
func (_NgmiTreasury *NgmiTreasuryTransactorSession) AddTier(is_new_tier_purchasable bool, new_tier_duration_days *big.Int, new_tier_prices []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.AddTier(&_NgmiTreasury.TransactOpts, is_new_tier_purchasable, new_tier_duration_days, new_tier_prices)
}

// DeactivateAllTiers is a paid mutator transaction binding the contract method 0xc5498885.
//
// Solidity: function deactivateAllTiers() returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) DeactivateAllTiers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "deactivateAllTiers")
}

// DeactivateAllTiers is a paid mutator transaction binding the contract method 0xc5498885.
//
// Solidity: function deactivateAllTiers() returns()
func (_NgmiTreasury *NgmiTreasurySession) DeactivateAllTiers() (*types.Transaction, error) {
	return _NgmiTreasury.Contract.DeactivateAllTiers(&_NgmiTreasury.TransactOpts)
}

// DeactivateAllTiers is a paid mutator transaction binding the contract method 0xc5498885.
//
// Solidity: function deactivateAllTiers() returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) DeactivateAllTiers() (*types.Transaction, error) {
	return _NgmiTreasury.Contract.DeactivateAllTiers(&_NgmiTreasury.TransactOpts)
}

// DeactivateTier is a paid mutator transaction binding the contract method 0x8bd684ee.
//
// Solidity: function deactivateTier(uint256 i_tier) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) DeactivateTier(opts *bind.TransactOpts, i_tier *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "deactivateTier", i_tier)
}

// DeactivateTier is a paid mutator transaction binding the contract method 0x8bd684ee.
//
// Solidity: function deactivateTier(uint256 i_tier) returns()
func (_NgmiTreasury *NgmiTreasurySession) DeactivateTier(i_tier *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.DeactivateTier(&_NgmiTreasury.TransactOpts, i_tier)
}

// DeactivateTier is a paid mutator transaction binding the contract method 0x8bd684ee.
//
// Solidity: function deactivateTier(uint256 i_tier) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) DeactivateTier(i_tier *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.DeactivateTier(&_NgmiTreasury.TransactOpts, i_tier)
}

// ForbidTokenForPayment is a paid mutator transaction binding the contract method 0xae7315f3.
//
// Solidity: function forbidTokenForPayment(address token_address) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) ForbidTokenForPayment(opts *bind.TransactOpts, token_address common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "forbidTokenForPayment", token_address)
}

// ForbidTokenForPayment is a paid mutator transaction binding the contract method 0xae7315f3.
//
// Solidity: function forbidTokenForPayment(address token_address) returns()
func (_NgmiTreasury *NgmiTreasurySession) ForbidTokenForPayment(token_address common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.ForbidTokenForPayment(&_NgmiTreasury.TransactOpts, token_address)
}

// ForbidTokenForPayment is a paid mutator transaction binding the contract method 0xae7315f3.
//
// Solidity: function forbidTokenForPayment(address token_address) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) ForbidTokenForPayment(token_address common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.ForbidTokenForPayment(&_NgmiTreasury.TransactOpts, token_address)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address account) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) Release(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "release", account)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address account) returns()
func (_NgmiTreasury *NgmiTreasurySession) Release(account common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.Release(&_NgmiTreasury.TransactOpts, account)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address account) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) Release(account common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.Release(&_NgmiTreasury.TransactOpts, account)
}

// Release0 is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address token, address account) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) Release0(opts *bind.TransactOpts, token common.Address, account common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "release0", token, account)
}

// Release0 is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address token, address account) returns()
func (_NgmiTreasury *NgmiTreasurySession) Release0(token common.Address, account common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.Release0(&_NgmiTreasury.TransactOpts, token, account)
}

// Release0 is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address token, address account) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) Release0(token common.Address, account common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.Release0(&_NgmiTreasury.TransactOpts, token, account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address removed_admin) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) RemoveAdmin(opts *bind.TransactOpts, removed_admin common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "removeAdmin", removed_admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address removed_admin) returns()
func (_NgmiTreasury *NgmiTreasurySession) RemoveAdmin(removed_admin common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.RemoveAdmin(&_NgmiTreasury.TransactOpts, removed_admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address removed_admin) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) RemoveAdmin(removed_admin common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.RemoveAdmin(&_NgmiTreasury.TransactOpts, removed_admin)
}

// RemoveAdmins is a paid mutator transaction binding the contract method 0x377e11e0.
//
// Solidity: function removeAdmins(address[] removed_admins) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) RemoveAdmins(opts *bind.TransactOpts, removed_admins []common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "removeAdmins", removed_admins)
}

// RemoveAdmins is a paid mutator transaction binding the contract method 0x377e11e0.
//
// Solidity: function removeAdmins(address[] removed_admins) returns()
func (_NgmiTreasury *NgmiTreasurySession) RemoveAdmins(removed_admins []common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.RemoveAdmins(&_NgmiTreasury.TransactOpts, removed_admins)
}

// RemoveAdmins is a paid mutator transaction binding the contract method 0x377e11e0.
//
// Solidity: function removeAdmins(address[] removed_admins) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) RemoveAdmins(removed_admins []common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.RemoveAdmins(&_NgmiTreasury.TransactOpts, removed_admins)
}

// SetAllConsideredTokensAcceptances is a paid mutator transaction binding the contract method 0xc6cab224.
//
// Solidity: function setAllConsideredTokensAcceptances(bool[] new_is_token_accepted_array) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) SetAllConsideredTokensAcceptances(opts *bind.TransactOpts, new_is_token_accepted_array []bool) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "setAllConsideredTokensAcceptances", new_is_token_accepted_array)
}

// SetAllConsideredTokensAcceptances is a paid mutator transaction binding the contract method 0xc6cab224.
//
// Solidity: function setAllConsideredTokensAcceptances(bool[] new_is_token_accepted_array) returns()
func (_NgmiTreasury *NgmiTreasurySession) SetAllConsideredTokensAcceptances(new_is_token_accepted_array []bool) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetAllConsideredTokensAcceptances(&_NgmiTreasury.TransactOpts, new_is_token_accepted_array)
}

// SetAllConsideredTokensAcceptances is a paid mutator transaction binding the contract method 0xc6cab224.
//
// Solidity: function setAllConsideredTokensAcceptances(bool[] new_is_token_accepted_array) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) SetAllConsideredTokensAcceptances(new_is_token_accepted_array []bool) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetAllConsideredTokensAcceptances(&_NgmiTreasury.TransactOpts, new_is_token_accepted_array)
}

// SetAllTierDurationsInDays is a paid mutator transaction binding the contract method 0xbeb9644d.
//
// Solidity: function setAllTierDurationsInDays(uint256[] new_tier_durations_days) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) SetAllTierDurationsInDays(opts *bind.TransactOpts, new_tier_durations_days []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "setAllTierDurationsInDays", new_tier_durations_days)
}

// SetAllTierDurationsInDays is a paid mutator transaction binding the contract method 0xbeb9644d.
//
// Solidity: function setAllTierDurationsInDays(uint256[] new_tier_durations_days) returns()
func (_NgmiTreasury *NgmiTreasurySession) SetAllTierDurationsInDays(new_tier_durations_days []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetAllTierDurationsInDays(&_NgmiTreasury.TransactOpts, new_tier_durations_days)
}

// SetAllTierDurationsInDays is a paid mutator transaction binding the contract method 0xbeb9644d.
//
// Solidity: function setAllTierDurationsInDays(uint256[] new_tier_durations_days) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) SetAllTierDurationsInDays(new_tier_durations_days []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetAllTierDurationsInDays(&_NgmiTreasury.TransactOpts, new_tier_durations_days)
}

// SetMultipleUserFirstSubscriptionTimestamps is a paid mutator transaction binding the contract method 0x99ea50b9.
//
// Solidity: function setMultipleUserFirstSubscriptionTimestamps(address[] users, uint256[] first_subscription_timestamps) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) SetMultipleUserFirstSubscriptionTimestamps(opts *bind.TransactOpts, users []common.Address, first_subscription_timestamps []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "setMultipleUserFirstSubscriptionTimestamps", users, first_subscription_timestamps)
}

// SetMultipleUserFirstSubscriptionTimestamps is a paid mutator transaction binding the contract method 0x99ea50b9.
//
// Solidity: function setMultipleUserFirstSubscriptionTimestamps(address[] users, uint256[] first_subscription_timestamps) returns()
func (_NgmiTreasury *NgmiTreasurySession) SetMultipleUserFirstSubscriptionTimestamps(users []common.Address, first_subscription_timestamps []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetMultipleUserFirstSubscriptionTimestamps(&_NgmiTreasury.TransactOpts, users, first_subscription_timestamps)
}

// SetMultipleUserFirstSubscriptionTimestamps is a paid mutator transaction binding the contract method 0x99ea50b9.
//
// Solidity: function setMultipleUserFirstSubscriptionTimestamps(address[] users, uint256[] first_subscription_timestamps) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) SetMultipleUserFirstSubscriptionTimestamps(users []common.Address, first_subscription_timestamps []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetMultipleUserFirstSubscriptionTimestamps(&_NgmiTreasury.TransactOpts, users, first_subscription_timestamps)
}

// SetMultipleUserLatestSubscriptionTimestamps is a paid mutator transaction binding the contract method 0x40b8f58b.
//
// Solidity: function setMultipleUserLatestSubscriptionTimestamps(address[] users, uint256[] latest_subscription_timestamps) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) SetMultipleUserLatestSubscriptionTimestamps(opts *bind.TransactOpts, users []common.Address, latest_subscription_timestamps []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "setMultipleUserLatestSubscriptionTimestamps", users, latest_subscription_timestamps)
}

// SetMultipleUserLatestSubscriptionTimestamps is a paid mutator transaction binding the contract method 0x40b8f58b.
//
// Solidity: function setMultipleUserLatestSubscriptionTimestamps(address[] users, uint256[] latest_subscription_timestamps) returns()
func (_NgmiTreasury *NgmiTreasurySession) SetMultipleUserLatestSubscriptionTimestamps(users []common.Address, latest_subscription_timestamps []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetMultipleUserLatestSubscriptionTimestamps(&_NgmiTreasury.TransactOpts, users, latest_subscription_timestamps)
}

// SetMultipleUserLatestSubscriptionTimestamps is a paid mutator transaction binding the contract method 0x40b8f58b.
//
// Solidity: function setMultipleUserLatestSubscriptionTimestamps(address[] users, uint256[] latest_subscription_timestamps) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) SetMultipleUserLatestSubscriptionTimestamps(users []common.Address, latest_subscription_timestamps []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetMultipleUserLatestSubscriptionTimestamps(&_NgmiTreasury.TransactOpts, users, latest_subscription_timestamps)
}

// SetMultipleUserTiers is a paid mutator transaction binding the contract method 0x5ceabc26.
//
// Solidity: function setMultipleUserTiers(address[] users, uint256[] new_i_tiers) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) SetMultipleUserTiers(opts *bind.TransactOpts, users []common.Address, new_i_tiers []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "setMultipleUserTiers", users, new_i_tiers)
}

// SetMultipleUserTiers is a paid mutator transaction binding the contract method 0x5ceabc26.
//
// Solidity: function setMultipleUserTiers(address[] users, uint256[] new_i_tiers) returns()
func (_NgmiTreasury *NgmiTreasurySession) SetMultipleUserTiers(users []common.Address, new_i_tiers []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetMultipleUserTiers(&_NgmiTreasury.TransactOpts, users, new_i_tiers)
}

// SetMultipleUserTiers is a paid mutator transaction binding the contract method 0x5ceabc26.
//
// Solidity: function setMultipleUserTiers(address[] users, uint256[] new_i_tiers) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) SetMultipleUserTiers(users []common.Address, new_i_tiers []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetMultipleUserTiers(&_NgmiTreasury.TransactOpts, users, new_i_tiers)
}

// SetNewPayees is a paid mutator transaction binding the contract method 0x91bfdd1a.
//
// Solidity: function setNewPayees(address[] payees, uint256[] shares_) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) SetNewPayees(opts *bind.TransactOpts, payees []common.Address, shares_ []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "setNewPayees", payees, shares_)
}

// SetNewPayees is a paid mutator transaction binding the contract method 0x91bfdd1a.
//
// Solidity: function setNewPayees(address[] payees, uint256[] shares_) returns()
func (_NgmiTreasury *NgmiTreasurySession) SetNewPayees(payees []common.Address, shares_ []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetNewPayees(&_NgmiTreasury.TransactOpts, payees, shares_)
}

// SetNewPayees is a paid mutator transaction binding the contract method 0x91bfdd1a.
//
// Solidity: function setNewPayees(address[] payees, uint256[] shares_) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) SetNewPayees(payees []common.Address, shares_ []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetNewPayees(&_NgmiTreasury.TransactOpts, payees, shares_)
}

// SetTierDurationInDays is a paid mutator transaction binding the contract method 0xec8ac0d8.
//
// Solidity: function setTierDurationInDays(uint256 i_tier, uint256 new_tier_duration_days) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) SetTierDurationInDays(opts *bind.TransactOpts, i_tier *big.Int, new_tier_duration_days *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "setTierDurationInDays", i_tier, new_tier_duration_days)
}

// SetTierDurationInDays is a paid mutator transaction binding the contract method 0xec8ac0d8.
//
// Solidity: function setTierDurationInDays(uint256 i_tier, uint256 new_tier_duration_days) returns()
func (_NgmiTreasury *NgmiTreasurySession) SetTierDurationInDays(i_tier *big.Int, new_tier_duration_days *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetTierDurationInDays(&_NgmiTreasury.TransactOpts, i_tier, new_tier_duration_days)
}

// SetTierDurationInDays is a paid mutator transaction binding the contract method 0xec8ac0d8.
//
// Solidity: function setTierDurationInDays(uint256 i_tier, uint256 new_tier_duration_days) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) SetTierDurationInDays(i_tier *big.Int, new_tier_duration_days *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetTierDurationInDays(&_NgmiTreasury.TransactOpts, i_tier, new_tier_duration_days)
}

// SetTierPriceInToken is a paid mutator transaction binding the contract method 0x5b7949cc.
//
// Solidity: function setTierPriceInToken(uint256 i_tier, address token_address, uint256 new_tier_price_token) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) SetTierPriceInToken(opts *bind.TransactOpts, i_tier *big.Int, token_address common.Address, new_tier_price_token *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "setTierPriceInToken", i_tier, token_address, new_tier_price_token)
}

// SetTierPriceInToken is a paid mutator transaction binding the contract method 0x5b7949cc.
//
// Solidity: function setTierPriceInToken(uint256 i_tier, address token_address, uint256 new_tier_price_token) returns()
func (_NgmiTreasury *NgmiTreasurySession) SetTierPriceInToken(i_tier *big.Int, token_address common.Address, new_tier_price_token *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetTierPriceInToken(&_NgmiTreasury.TransactOpts, i_tier, token_address, new_tier_price_token)
}

// SetTierPriceInToken is a paid mutator transaction binding the contract method 0x5b7949cc.
//
// Solidity: function setTierPriceInToken(uint256 i_tier, address token_address, uint256 new_tier_price_token) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) SetTierPriceInToken(i_tier *big.Int, token_address common.Address, new_tier_price_token *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetTierPriceInToken(&_NgmiTreasury.TransactOpts, i_tier, token_address, new_tier_price_token)
}

// SetTierPricesForAllConsideredTokens is a paid mutator transaction binding the contract method 0xa81963a8.
//
// Solidity: function setTierPricesForAllConsideredTokens(uint256 i_tier, uint256[] new_tier_price_in_all_tokens) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) SetTierPricesForAllConsideredTokens(opts *bind.TransactOpts, i_tier *big.Int, new_tier_price_in_all_tokens []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "setTierPricesForAllConsideredTokens", i_tier, new_tier_price_in_all_tokens)
}

// SetTierPricesForAllConsideredTokens is a paid mutator transaction binding the contract method 0xa81963a8.
//
// Solidity: function setTierPricesForAllConsideredTokens(uint256 i_tier, uint256[] new_tier_price_in_all_tokens) returns()
func (_NgmiTreasury *NgmiTreasurySession) SetTierPricesForAllConsideredTokens(i_tier *big.Int, new_tier_price_in_all_tokens []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetTierPricesForAllConsideredTokens(&_NgmiTreasury.TransactOpts, i_tier, new_tier_price_in_all_tokens)
}

// SetTierPricesForAllConsideredTokens is a paid mutator transaction binding the contract method 0xa81963a8.
//
// Solidity: function setTierPricesForAllConsideredTokens(uint256 i_tier, uint256[] new_tier_price_in_all_tokens) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) SetTierPricesForAllConsideredTokens(i_tier *big.Int, new_tier_price_in_all_tokens []*big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetTierPricesForAllConsideredTokens(&_NgmiTreasury.TransactOpts, i_tier, new_tier_price_in_all_tokens)
}

// SetUserFirstSubscriptionTimestamp is a paid mutator transaction binding the contract method 0xe8735f29.
//
// Solidity: function setUserFirstSubscriptionTimestamp(address user, uint256 first_subscription_timestamp) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) SetUserFirstSubscriptionTimestamp(opts *bind.TransactOpts, user common.Address, first_subscription_timestamp *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "setUserFirstSubscriptionTimestamp", user, first_subscription_timestamp)
}

// SetUserFirstSubscriptionTimestamp is a paid mutator transaction binding the contract method 0xe8735f29.
//
// Solidity: function setUserFirstSubscriptionTimestamp(address user, uint256 first_subscription_timestamp) returns()
func (_NgmiTreasury *NgmiTreasurySession) SetUserFirstSubscriptionTimestamp(user common.Address, first_subscription_timestamp *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetUserFirstSubscriptionTimestamp(&_NgmiTreasury.TransactOpts, user, first_subscription_timestamp)
}

// SetUserFirstSubscriptionTimestamp is a paid mutator transaction binding the contract method 0xe8735f29.
//
// Solidity: function setUserFirstSubscriptionTimestamp(address user, uint256 first_subscription_timestamp) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) SetUserFirstSubscriptionTimestamp(user common.Address, first_subscription_timestamp *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetUserFirstSubscriptionTimestamp(&_NgmiTreasury.TransactOpts, user, first_subscription_timestamp)
}

// SetUserLatestSubscriptionTimestamp is a paid mutator transaction binding the contract method 0xfe6cf8d3.
//
// Solidity: function setUserLatestSubscriptionTimestamp(address user, uint256 latest_subscription_timestamp) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) SetUserLatestSubscriptionTimestamp(opts *bind.TransactOpts, user common.Address, latest_subscription_timestamp *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "setUserLatestSubscriptionTimestamp", user, latest_subscription_timestamp)
}

// SetUserLatestSubscriptionTimestamp is a paid mutator transaction binding the contract method 0xfe6cf8d3.
//
// Solidity: function setUserLatestSubscriptionTimestamp(address user, uint256 latest_subscription_timestamp) returns()
func (_NgmiTreasury *NgmiTreasurySession) SetUserLatestSubscriptionTimestamp(user common.Address, latest_subscription_timestamp *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetUserLatestSubscriptionTimestamp(&_NgmiTreasury.TransactOpts, user, latest_subscription_timestamp)
}

// SetUserLatestSubscriptionTimestamp is a paid mutator transaction binding the contract method 0xfe6cf8d3.
//
// Solidity: function setUserLatestSubscriptionTimestamp(address user, uint256 latest_subscription_timestamp) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) SetUserLatestSubscriptionTimestamp(user common.Address, latest_subscription_timestamp *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetUserLatestSubscriptionTimestamp(&_NgmiTreasury.TransactOpts, user, latest_subscription_timestamp)
}

// SetUserTier is a paid mutator transaction binding the contract method 0x498c35fd.
//
// Solidity: function setUserTier(address user, uint256 new_i_tier) returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) SetUserTier(opts *bind.TransactOpts, user common.Address, new_i_tier *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "setUserTier", user, new_i_tier)
}

// SetUserTier is a paid mutator transaction binding the contract method 0x498c35fd.
//
// Solidity: function setUserTier(address user, uint256 new_i_tier) returns()
func (_NgmiTreasury *NgmiTreasurySession) SetUserTier(user common.Address, new_i_tier *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetUserTier(&_NgmiTreasury.TransactOpts, user, new_i_tier)
}

// SetUserTier is a paid mutator transaction binding the contract method 0x498c35fd.
//
// Solidity: function setUserTier(address user, uint256 new_i_tier) returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) SetUserTier(user common.Address, new_i_tier *big.Int) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.SetUserTier(&_NgmiTreasury.TransactOpts, user, new_i_tier)
}

// Subscribe is a paid mutator transaction binding the contract method 0x59e6951d.
//
// Solidity: function subscribe(uint256 i_tier, address token_address) payable returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) Subscribe(opts *bind.TransactOpts, i_tier *big.Int, token_address common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "subscribe", i_tier, token_address)
}

// Subscribe is a paid mutator transaction binding the contract method 0x59e6951d.
//
// Solidity: function subscribe(uint256 i_tier, address token_address) payable returns()
func (_NgmiTreasury *NgmiTreasurySession) Subscribe(i_tier *big.Int, token_address common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.Subscribe(&_NgmiTreasury.TransactOpts, i_tier, token_address)
}

// Subscribe is a paid mutator transaction binding the contract method 0x59e6951d.
//
// Solidity: function subscribe(uint256 i_tier, address token_address) payable returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) Subscribe(i_tier *big.Int, token_address common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.Subscribe(&_NgmiTreasury.TransactOpts, i_tier, token_address)
}

// Upgrade is a paid mutator transaction binding the contract method 0x028f4e47.
//
// Solidity: function upgrade(uint256 i_new_tier, address token_address) payable returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) Upgrade(opts *bind.TransactOpts, i_new_tier *big.Int, token_address common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.contract.Transact(opts, "upgrade", i_new_tier, token_address)
}

// Upgrade is a paid mutator transaction binding the contract method 0x028f4e47.
//
// Solidity: function upgrade(uint256 i_new_tier, address token_address) payable returns()
func (_NgmiTreasury *NgmiTreasurySession) Upgrade(i_new_tier *big.Int, token_address common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.Upgrade(&_NgmiTreasury.TransactOpts, i_new_tier, token_address)
}

// Upgrade is a paid mutator transaction binding the contract method 0x028f4e47.
//
// Solidity: function upgrade(uint256 i_new_tier, address token_address) payable returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) Upgrade(i_new_tier *big.Int, token_address common.Address) (*types.Transaction, error) {
	return _NgmiTreasury.Contract.Upgrade(&_NgmiTreasury.TransactOpts, i_new_tier, token_address)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NgmiTreasury *NgmiTreasuryTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NgmiTreasury.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NgmiTreasury *NgmiTreasurySession) Receive() (*types.Transaction, error) {
	return _NgmiTreasury.Contract.Receive(&_NgmiTreasury.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NgmiTreasury *NgmiTreasuryTransactorSession) Receive() (*types.Transaction, error) {
	return _NgmiTreasury.Contract.Receive(&_NgmiTreasury.TransactOpts)
}

// NgmiTreasuryERC20PaymentReleasedIterator is returned from FilterERC20PaymentReleased and is used to iterate over the raw logs and unpacked data for ERC20PaymentReleased events raised by the NgmiTreasury contract.
type NgmiTreasuryERC20PaymentReleasedIterator struct {
	Event *NgmiTreasuryERC20PaymentReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NgmiTreasuryERC20PaymentReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NgmiTreasuryERC20PaymentReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NgmiTreasuryERC20PaymentReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NgmiTreasuryERC20PaymentReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NgmiTreasuryERC20PaymentReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NgmiTreasuryERC20PaymentReleased represents a ERC20PaymentReleased event raised by the NgmiTreasury contract.
type NgmiTreasuryERC20PaymentReleased struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC20PaymentReleased is a free log retrieval operation binding the contract event 0x3be5b7a71e84ed12875d241991c70855ac5817d847039e17a9d895c1ceb0f18a.
//
// Solidity: event ERC20PaymentReleased(address indexed token, address to, uint256 amount)
func (_NgmiTreasury *NgmiTreasuryFilterer) FilterERC20PaymentReleased(opts *bind.FilterOpts, token []common.Address) (*NgmiTreasuryERC20PaymentReleasedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NgmiTreasury.contract.FilterLogs(opts, "ERC20PaymentReleased", tokenRule)
	if err != nil {
		return nil, err
	}
	return &NgmiTreasuryERC20PaymentReleasedIterator{contract: _NgmiTreasury.contract, event: "ERC20PaymentReleased", logs: logs, sub: sub}, nil
}

// WatchERC20PaymentReleased is a free log subscription operation binding the contract event 0x3be5b7a71e84ed12875d241991c70855ac5817d847039e17a9d895c1ceb0f18a.
//
// Solidity: event ERC20PaymentReleased(address indexed token, address to, uint256 amount)
func (_NgmiTreasury *NgmiTreasuryFilterer) WatchERC20PaymentReleased(opts *bind.WatchOpts, sink chan<- *NgmiTreasuryERC20PaymentReleased, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _NgmiTreasury.contract.WatchLogs(opts, "ERC20PaymentReleased", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NgmiTreasuryERC20PaymentReleased)
				if err := _NgmiTreasury.contract.UnpackLog(event, "ERC20PaymentReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseERC20PaymentReleased is a log parse operation binding the contract event 0x3be5b7a71e84ed12875d241991c70855ac5817d847039e17a9d895c1ceb0f18a.
//
// Solidity: event ERC20PaymentReleased(address indexed token, address to, uint256 amount)
func (_NgmiTreasury *NgmiTreasuryFilterer) ParseERC20PaymentReleased(log types.Log) (*NgmiTreasuryERC20PaymentReleased, error) {
	event := new(NgmiTreasuryERC20PaymentReleased)
	if err := _NgmiTreasury.contract.UnpackLog(event, "ERC20PaymentReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NgmiTreasuryNewSubscriptionIterator is returned from FilterNewSubscription and is used to iterate over the raw logs and unpacked data for NewSubscription events raised by the NgmiTreasury contract.
type NgmiTreasuryNewSubscriptionIterator struct {
	Event *NgmiTreasuryNewSubscription // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NgmiTreasuryNewSubscriptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NgmiTreasuryNewSubscription)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NgmiTreasuryNewSubscription)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NgmiTreasuryNewSubscriptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NgmiTreasuryNewSubscriptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NgmiTreasuryNewSubscription represents a NewSubscription event raised by the NgmiTreasury contract.
type NgmiTreasuryNewSubscription struct {
	Timestamp        *big.Int
	Subscriber       common.Address
	ITier            *big.Int
	TierDurationDays *big.Int
	TokenAddress     common.Address
	TierPriceInToken *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewSubscription is a free log retrieval operation binding the contract event 0x1b43a21695c02f3b024f6efacda3d4611eed255b95a885def65416a7aa902580.
//
// Solidity: event NewSubscription(uint256 timestamp, address subscriber, uint256 i_tier, uint256 tier_duration_days, address token_address, uint256 tier_price_in_token)
func (_NgmiTreasury *NgmiTreasuryFilterer) FilterNewSubscription(opts *bind.FilterOpts) (*NgmiTreasuryNewSubscriptionIterator, error) {

	logs, sub, err := _NgmiTreasury.contract.FilterLogs(opts, "NewSubscription")
	if err != nil {
		return nil, err
	}
	return &NgmiTreasuryNewSubscriptionIterator{contract: _NgmiTreasury.contract, event: "NewSubscription", logs: logs, sub: sub}, nil
}

// WatchNewSubscription is a free log subscription operation binding the contract event 0x1b43a21695c02f3b024f6efacda3d4611eed255b95a885def65416a7aa902580.
//
// Solidity: event NewSubscription(uint256 timestamp, address subscriber, uint256 i_tier, uint256 tier_duration_days, address token_address, uint256 tier_price_in_token)
func (_NgmiTreasury *NgmiTreasuryFilterer) WatchNewSubscription(opts *bind.WatchOpts, sink chan<- *NgmiTreasuryNewSubscription) (event.Subscription, error) {

	logs, sub, err := _NgmiTreasury.contract.WatchLogs(opts, "NewSubscription")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NgmiTreasuryNewSubscription)
				if err := _NgmiTreasury.contract.UnpackLog(event, "NewSubscription", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewSubscription is a log parse operation binding the contract event 0x1b43a21695c02f3b024f6efacda3d4611eed255b95a885def65416a7aa902580.
//
// Solidity: event NewSubscription(uint256 timestamp, address subscriber, uint256 i_tier, uint256 tier_duration_days, address token_address, uint256 tier_price_in_token)
func (_NgmiTreasury *NgmiTreasuryFilterer) ParseNewSubscription(log types.Log) (*NgmiTreasuryNewSubscription, error) {
	event := new(NgmiTreasuryNewSubscription)
	if err := _NgmiTreasury.contract.UnpackLog(event, "NewSubscription", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NgmiTreasuryPayeeAddedIterator is returned from FilterPayeeAdded and is used to iterate over the raw logs and unpacked data for PayeeAdded events raised by the NgmiTreasury contract.
type NgmiTreasuryPayeeAddedIterator struct {
	Event *NgmiTreasuryPayeeAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NgmiTreasuryPayeeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NgmiTreasuryPayeeAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NgmiTreasuryPayeeAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NgmiTreasuryPayeeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NgmiTreasuryPayeeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NgmiTreasuryPayeeAdded represents a PayeeAdded event raised by the NgmiTreasury contract.
type NgmiTreasuryPayeeAdded struct {
	Account common.Address
	Shares  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPayeeAdded is a free log retrieval operation binding the contract event 0x40c340f65e17194d14ddddb073d3c9f888e3cb52b5aae0c6c7706b4fbc905fac.
//
// Solidity: event PayeeAdded(address account, uint256 shares)
func (_NgmiTreasury *NgmiTreasuryFilterer) FilterPayeeAdded(opts *bind.FilterOpts) (*NgmiTreasuryPayeeAddedIterator, error) {

	logs, sub, err := _NgmiTreasury.contract.FilterLogs(opts, "PayeeAdded")
	if err != nil {
		return nil, err
	}
	return &NgmiTreasuryPayeeAddedIterator{contract: _NgmiTreasury.contract, event: "PayeeAdded", logs: logs, sub: sub}, nil
}

// WatchPayeeAdded is a free log subscription operation binding the contract event 0x40c340f65e17194d14ddddb073d3c9f888e3cb52b5aae0c6c7706b4fbc905fac.
//
// Solidity: event PayeeAdded(address account, uint256 shares)
func (_NgmiTreasury *NgmiTreasuryFilterer) WatchPayeeAdded(opts *bind.WatchOpts, sink chan<- *NgmiTreasuryPayeeAdded) (event.Subscription, error) {

	logs, sub, err := _NgmiTreasury.contract.WatchLogs(opts, "PayeeAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NgmiTreasuryPayeeAdded)
				if err := _NgmiTreasury.contract.UnpackLog(event, "PayeeAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePayeeAdded is a log parse operation binding the contract event 0x40c340f65e17194d14ddddb073d3c9f888e3cb52b5aae0c6c7706b4fbc905fac.
//
// Solidity: event PayeeAdded(address account, uint256 shares)
func (_NgmiTreasury *NgmiTreasuryFilterer) ParsePayeeAdded(log types.Log) (*NgmiTreasuryPayeeAdded, error) {
	event := new(NgmiTreasuryPayeeAdded)
	if err := _NgmiTreasury.contract.UnpackLog(event, "PayeeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NgmiTreasuryPaymentReceivedIterator is returned from FilterPaymentReceived and is used to iterate over the raw logs and unpacked data for PaymentReceived events raised by the NgmiTreasury contract.
type NgmiTreasuryPaymentReceivedIterator struct {
	Event *NgmiTreasuryPaymentReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NgmiTreasuryPaymentReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NgmiTreasuryPaymentReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NgmiTreasuryPaymentReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NgmiTreasuryPaymentReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NgmiTreasuryPaymentReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NgmiTreasuryPaymentReceived represents a PaymentReceived event raised by the NgmiTreasury contract.
type NgmiTreasuryPaymentReceived struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentReceived is a free log retrieval operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_NgmiTreasury *NgmiTreasuryFilterer) FilterPaymentReceived(opts *bind.FilterOpts) (*NgmiTreasuryPaymentReceivedIterator, error) {

	logs, sub, err := _NgmiTreasury.contract.FilterLogs(opts, "PaymentReceived")
	if err != nil {
		return nil, err
	}
	return &NgmiTreasuryPaymentReceivedIterator{contract: _NgmiTreasury.contract, event: "PaymentReceived", logs: logs, sub: sub}, nil
}

// WatchPaymentReceived is a free log subscription operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_NgmiTreasury *NgmiTreasuryFilterer) WatchPaymentReceived(opts *bind.WatchOpts, sink chan<- *NgmiTreasuryPaymentReceived) (event.Subscription, error) {

	logs, sub, err := _NgmiTreasury.contract.WatchLogs(opts, "PaymentReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NgmiTreasuryPaymentReceived)
				if err := _NgmiTreasury.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentReceived is a log parse operation binding the contract event 0x6ef95f06320e7a25a04a175ca677b7052bdd97131872c2192525a629f51be770.
//
// Solidity: event PaymentReceived(address from, uint256 amount)
func (_NgmiTreasury *NgmiTreasuryFilterer) ParsePaymentReceived(log types.Log) (*NgmiTreasuryPaymentReceived, error) {
	event := new(NgmiTreasuryPaymentReceived)
	if err := _NgmiTreasury.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NgmiTreasuryPaymentReleasedIterator is returned from FilterPaymentReleased and is used to iterate over the raw logs and unpacked data for PaymentReleased events raised by the NgmiTreasury contract.
type NgmiTreasuryPaymentReleasedIterator struct {
	Event *NgmiTreasuryPaymentReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NgmiTreasuryPaymentReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NgmiTreasuryPaymentReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NgmiTreasuryPaymentReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NgmiTreasuryPaymentReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NgmiTreasuryPaymentReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NgmiTreasuryPaymentReleased represents a PaymentReleased event raised by the NgmiTreasury contract.
type NgmiTreasuryPaymentReleased struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaymentReleased is a free log retrieval operation binding the contract event 0xdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056.
//
// Solidity: event PaymentReleased(address to, uint256 amount)
func (_NgmiTreasury *NgmiTreasuryFilterer) FilterPaymentReleased(opts *bind.FilterOpts) (*NgmiTreasuryPaymentReleasedIterator, error) {

	logs, sub, err := _NgmiTreasury.contract.FilterLogs(opts, "PaymentReleased")
	if err != nil {
		return nil, err
	}
	return &NgmiTreasuryPaymentReleasedIterator{contract: _NgmiTreasury.contract, event: "PaymentReleased", logs: logs, sub: sub}, nil
}

// WatchPaymentReleased is a free log subscription operation binding the contract event 0xdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056.
//
// Solidity: event PaymentReleased(address to, uint256 amount)
func (_NgmiTreasury *NgmiTreasuryFilterer) WatchPaymentReleased(opts *bind.WatchOpts, sink chan<- *NgmiTreasuryPaymentReleased) (event.Subscription, error) {

	logs, sub, err := _NgmiTreasury.contract.WatchLogs(opts, "PaymentReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NgmiTreasuryPaymentReleased)
				if err := _NgmiTreasury.contract.UnpackLog(event, "PaymentReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaymentReleased is a log parse operation binding the contract event 0xdf20fd1e76bc69d672e4814fafb2c449bba3a5369d8359adf9e05e6fde87b056.
//
// Solidity: event PaymentReleased(address to, uint256 amount)
func (_NgmiTreasury *NgmiTreasuryFilterer) ParsePaymentReleased(log types.Log) (*NgmiTreasuryPaymentReleased, error) {
	event := new(NgmiTreasuryPaymentReleased)
	if err := _NgmiTreasury.contract.UnpackLog(event, "PaymentReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NgmiTreasuryStateResetIterator is returned from FilterStateReset and is used to iterate over the raw logs and unpacked data for StateReset events raised by the NgmiTreasury contract.
type NgmiTreasuryStateResetIterator struct {
	Event *NgmiTreasuryStateReset // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NgmiTreasuryStateResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NgmiTreasuryStateReset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NgmiTreasuryStateReset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NgmiTreasuryStateResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NgmiTreasuryStateResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NgmiTreasuryStateReset represents a StateReset event raised by the NgmiTreasury contract.
type NgmiTreasuryStateReset struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStateReset is a free log retrieval operation binding the contract event 0x5a2a4a53daceb14c3d178e92f0da1bf67a87465998a163d815182375fea5c263.
//
// Solidity: event StateReset()
func (_NgmiTreasury *NgmiTreasuryFilterer) FilterStateReset(opts *bind.FilterOpts) (*NgmiTreasuryStateResetIterator, error) {

	logs, sub, err := _NgmiTreasury.contract.FilterLogs(opts, "StateReset")
	if err != nil {
		return nil, err
	}
	return &NgmiTreasuryStateResetIterator{contract: _NgmiTreasury.contract, event: "StateReset", logs: logs, sub: sub}, nil
}

// WatchStateReset is a free log subscription operation binding the contract event 0x5a2a4a53daceb14c3d178e92f0da1bf67a87465998a163d815182375fea5c263.
//
// Solidity: event StateReset()
func (_NgmiTreasury *NgmiTreasuryFilterer) WatchStateReset(opts *bind.WatchOpts, sink chan<- *NgmiTreasuryStateReset) (event.Subscription, error) {

	logs, sub, err := _NgmiTreasury.contract.WatchLogs(opts, "StateReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NgmiTreasuryStateReset)
				if err := _NgmiTreasury.contract.UnpackLog(event, "StateReset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStateReset is a log parse operation binding the contract event 0x5a2a4a53daceb14c3d178e92f0da1bf67a87465998a163d815182375fea5c263.
//
// Solidity: event StateReset()
func (_NgmiTreasury *NgmiTreasuryFilterer) ParseStateReset(log types.Log) (*NgmiTreasuryStateReset, error) {
	event := new(NgmiTreasuryStateReset)
	if err := _NgmiTreasury.contract.UnpackLog(event, "StateReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
