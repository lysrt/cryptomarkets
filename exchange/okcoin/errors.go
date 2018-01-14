package okcoin

// HTTP ERROR CODE 403 	Too many requests, IP is shielded.
// Request Timed Out 	Too many requests, IP is shielded.

var errorCodes = map[int]string{
	// Error Code (Spot Market)
	10000: "Required field, can not be null",
	10001: "Request frequency too high to exceed the limit allowed",
	10002: "System error",
	10003: "Not in reqest list, please try again later",
	10004: "This IP is not allowed to access",
	10005: "'SecretKey' does not exist",
	10006: "'Api_key' does not exist",
	10007: "Signature does not match",
	10008: "Illegal parameter",
	10009: "Order does not exist",
	10010: "Insufficient funds",
	10011: "Amount too low",
	10012: "Only btc_usd ltc_usd supported",
	10013: "Only support https request",
	10014: "Order price must be between 0 and 1,000,000",
	10015: "Order price differs from current market price too much",
	10016: "Insufficient coins balance",
	10017: "API authorization error",
	10018: "borrow amount less than lower limit [usd:100,btc:0.1,ltc:1]",
	10019: "loan agreement not checked",
	10020: "rate cannot exceed 1%",
	10021: "rate cannot less than 0.01%",
	10023: "fail to get latest ticker",
	10024: "balance not sufficient",
	10025: "quota is full, cannot borrow temporarily",
	10026: "Loan (including reserved loan) and margin cannot be withdrawn",
	10027: "Cannot withdraw within 24 hrs of authentication information modification",
	10028: "Withdrawal amount exceeds daily limit",
	10029: "Account has unpaid loan, please cancel/pay off the loan before withdraw",
	10031: "Deposits can only be withdrawn after 6 confirmations",
	10032: "Please enabled phone/google authenticator",
	10033: "Fee higher than maximum network transaction fee",
	10034: "Fee lower than minimum network transaction fee",
	10035: "Insufficient BTC/LTC",
	10036: "Withdrawal amount too low",
	10037: "Trade password not set",
	10040: "Withdrawal cancellation fails",
	10041: "Withdrawal address not exist or approved",
	10042: "Admin password error",
	10043: "Account equity error, withdrawal failure",
	10044: "fail to cancel borrowing order",
	10047: "this function is disabled for sub-account",
	10048: "withdrawal information does not exist",
	10049: "User can not have more than 50 unfilled small orders (amount<0.15BTC)",
	10050: "can't cancel more than once",
	10052: "not allowed to withdraw",
	10064: "after a USD deposit, that portion of assets will not be withdrawable for the next 48 hours",
	10100: "User account frozen",
	10101: "order type is wrong",
	10102: "incorrect ID",
	10103: "the private otc order's key incorrect",
	10216: "Non-available API",
	1002:  "The transaction amount exceed the balance",
	1003:  "The transaction amount is less than the minimum requirement",
	1004:  "The transaction amount is less than 0",
	1007:  "No trading market information",
	1008:  "No latest market information",
	1009:  "No order",
	1010:  "Different user of the cancelled order and the original order",
	1011:  "No documented user",
	1013:  "No order type",
	1014:  "No login",
	1015:  "No market depth information",
	1017:  "Date error",
	1018:  "Order failed",
	1019:  "Undo order failed",
	1024:  "Currency does not exist",
	1025:  "No chart type",
	1026:  "No base currency quantity",
	1027:  "Incorrect parameter may exceeded limits",
	1028:  "Reserved decimal failed",
	1029:  "Preparing",
	1030:  "Account has margin and futures, transactions can not be processed",
	1031:  "Insufficient Transferring Balance",
	1032:  "Transferring Not Allowed",
	1035:  "Password incorrect",
	1036:  "Google Verification code Invalid",
	1037:  "Google Verification code incorrect",
	1038:  "Google Verification replicated",
	1039:  "Message Verification Input exceed the limit",
	1040:  "Message Verification invalid",
	1041:  "Message Verification incorrect",
	1042:  "Wrong Google Verification Input exceed the limit",
	1043:  "Login password cannot be same as the trading password",
	1044:  "Old password incorrect",
	1045:  "2nd Verification Needed",
	1046:  "Please input old password",
	1048:  "Account Blocked",
	1201:  "Account Deleted at 00: 00",
	1202:  "Account Not Exist",
	1203:  "Insufficient Balance",
	1204:  "Invalid currency",
	1205:  "Invalid Account",
	1206:  "Cash Withdrawal Blocked",
	1207:  "Transfer Not Support",
	1208:  "No designated account",
	1209:  "Invalid api",

	// Error Code (Contract API)
	20001: "User does not exist",
	20002: "Account frozen",
	20003: "Account frozen due to liquidation",
	20004: "Contract account frozen",
	20005: "User contract account does not exist",
	20006: "Required field missing",
	20007: "Illegal parameter",
	20008: "Contract account balance is too low",
	20009: "Contract status error",
	20010: "Risk rate ratio does not exist",
	20011: "Risk rate lower than 90% before opening position",
	20012: "Risk rate lower than 90% after opening position",
	20013: "Temporally no counter party price",
	20014: "System error",
	20015: "Order does not exist",
	20016: "Close amount bigger than your open positions",
	20017: "Not authorized/illegal operation",
	20018: "Order price cannot be more than 103% or less than 97% of the previous minute price",
	20019: "IP restricted from accessing the resource",
	20020: "secretKey does not exist",
	20021: "Index information does not exist",
	20022: "Wrong API interface (Cross margin mode shall call cross margin API, fixed margin mode shall call fixed margin API)",
	20023: "Account in fixed-margin mode",
	20024: "Signature does not match",
	20025: "Leverage rate error",
	20026: "API Permission Error",
	20027: "no transaction record",
	20028: "no such contract",
	20029: "Amount is large than available funds",
	20030: "Account still has debts",
	20038: "Due to regulation, this function is not availavle in the country/region your currently reside in.",
	20049: "Request frequency too high",
}
