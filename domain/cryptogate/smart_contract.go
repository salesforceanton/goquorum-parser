package cryptogate

import (
	"github.com/salesforceanton/goquorum-parser/services/backend/cryptogate/account_manager_abi"
	"github.com/salesforceanton/goquorum-parser/services/backend/cryptogate/node_manager_abi"
	"github.com/salesforceanton/goquorum-parser/services/backend/cryptogate/org_manager_abi"
	"github.com/salesforceanton/goquorum-parser/services/backend/cryptogate/permissions_impl_abi"
	"github.com/salesforceanton/goquorum-parser/services/backend/cryptogate/role_manager_abi"
	"github.com/salesforceanton/goquorum-parser/services/backend/cryptogate/voter_manager_abi"
)

type SmartContract struct {
	Address string
	Type    SmartContractType
}

type SmartContractType string

const (
	SmartContractTypePermissionImpl SmartContractType = "PERMISSIONS_IMPL"
	SmartContractTypeOrgManager     SmartContractType = "ORG_MANAGER"
	SmartContractTypeAccountManager SmartContractType = "ACCOUNT_MANAGER"
	SmartContractTypeNodeManager    SmartContractType = "NODE_MANAGER"
	SmartContractTypeRoleManager    SmartContractType = "ROLE_MANAGER"
	SmartContractTypeVoterManager   SmartContractType = "VOTER_MANAGER"
)

func GetABI(contractType SmartContractType) string {
	switch contractType {
	case SmartContractTypePermissionImpl:
		return permissions_impl_abi.PermissionsImplAbiMetaData.ABI
	case SmartContractTypeOrgManager:
		return org_manager_abi.OrgManagerAbiMetaData.ABI
	case SmartContractTypeAccountManager:
		return account_manager_abi.AccountManagerAbiMetaData.ABI
	case SmartContractTypeNodeManager:
		return node_manager_abi.NodeManagerAbiMetaData.ABI
	case SmartContractTypeRoleManager:
		return role_manager_abi.RoleManagerAbiMetaData.ABI
	case SmartContractTypeVoterManager:
		return voter_manager_abi.VoterManagerAbiMetaData.ABI
	default:
		return ""
	}
}
