// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeConformancePackTemplateValidationException for service response error code
	// "ConformancePackTemplateValidationException".
	//
	// You have specified a template that is not valid or supported.
	ErrCodeConformancePackTemplateValidationException = "ConformancePackTemplateValidationException"

	// ErrCodeInsufficientDeliveryPolicyException for service response error code
	// "InsufficientDeliveryPolicyException".
	//
	// Your Amazon S3 bucket policy does not permit AWS Config to write to it.
	ErrCodeInsufficientDeliveryPolicyException = "InsufficientDeliveryPolicyException"

	// ErrCodeInsufficientPermissionsException for service response error code
	// "InsufficientPermissionsException".
	//
	// Indicates one of the following errors:
	//
	//    * For PutConfigRule, the rule cannot be created because the IAM role assigned
	//    to AWS Config lacks permissions to perform the config:Put* action.
	//
	//    * For PutConfigRule, the AWS Lambda function cannot be invoked. Check
	//    the function ARN, and check the function's permissions.
	//
	//    * For PutOrganizationConfigRule, organization config rule cannot be created
	//    because you do not have permissions to call IAM GetRole action or create
	//    a service linked role.
	//
	//    * For PutConformancePack and PutOrganizationConformancePack, a conformance
	//    pack cannot be created because you do not have permissions: To call IAM
	//    GetRole action or create a service linked role. To read Amazon S3 bucket.
	ErrCodeInsufficientPermissionsException = "InsufficientPermissionsException"

	// ErrCodeInvalidConfigurationRecorderNameException for service response error code
	// "InvalidConfigurationRecorderNameException".
	//
	// You have provided a configuration recorder name that is not valid.
	ErrCodeInvalidConfigurationRecorderNameException = "InvalidConfigurationRecorderNameException"

	// ErrCodeInvalidDeliveryChannelNameException for service response error code
	// "InvalidDeliveryChannelNameException".
	//
	// The specified delivery channel name is not valid.
	ErrCodeInvalidDeliveryChannelNameException = "InvalidDeliveryChannelNameException"

	// ErrCodeInvalidExpressionException for service response error code
	// "InvalidExpressionException".
	//
	// The syntax of the query is incorrect.
	ErrCodeInvalidExpressionException = "InvalidExpressionException"

	// ErrCodeInvalidLimitException for service response error code
	// "InvalidLimitException".
	//
	// The specified limit is outside the allowable range.
	ErrCodeInvalidLimitException = "InvalidLimitException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The specified next token is invalid. Specify the nextToken string that was
	// returned in the previous response to get the next page of results.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidParameterValueException for service response error code
	// "InvalidParameterValueException".
	//
	// One or more of the specified parameters are invalid. Verify that your parameters
	// are valid and try again.
	ErrCodeInvalidParameterValueException = "InvalidParameterValueException"

	// ErrCodeInvalidRecordingGroupException for service response error code
	// "InvalidRecordingGroupException".
	//
	// AWS Config throws an exception if the recording group does not contain a
	// valid list of resource types. Invalid values might also be incorrectly formatted.
	ErrCodeInvalidRecordingGroupException = "InvalidRecordingGroupException"

	// ErrCodeInvalidResultTokenException for service response error code
	// "InvalidResultTokenException".
	//
	// The specified ResultToken is invalid.
	ErrCodeInvalidResultTokenException = "InvalidResultTokenException"

	// ErrCodeInvalidRoleException for service response error code
	// "InvalidRoleException".
	//
	// You have provided a null or empty role ARN.
	ErrCodeInvalidRoleException = "InvalidRoleException"

	// ErrCodeInvalidS3KeyPrefixException for service response error code
	// "InvalidS3KeyPrefixException".
	//
	// The specified Amazon S3 key prefix is not valid.
	ErrCodeInvalidS3KeyPrefixException = "InvalidS3KeyPrefixException"

	// ErrCodeInvalidSNSTopicARNException for service response error code
	// "InvalidSNSTopicARNException".
	//
	// The specified Amazon SNS topic does not exist.
	ErrCodeInvalidSNSTopicARNException = "InvalidSNSTopicARNException"

	// ErrCodeInvalidTimeRangeException for service response error code
	// "InvalidTimeRangeException".
	//
	// The specified time range is not valid. The earlier time is not chronologically
	// before the later time.
	ErrCodeInvalidTimeRangeException = "InvalidTimeRangeException"

	// ErrCodeLastDeliveryChannelDeleteFailedException for service response error code
	// "LastDeliveryChannelDeleteFailedException".
	//
	// You cannot delete the delivery channel you specified because the configuration
	// recorder is running.
	ErrCodeLastDeliveryChannelDeleteFailedException = "LastDeliveryChannelDeleteFailedException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// For StartConfigRulesEvaluation API, this exception is thrown if an evaluation
	// is in progress or if you call the StartConfigRulesEvaluation API more than
	// once per minute.
	//
	// For PutConfigurationAggregator API, this exception is thrown if the number
	// of accounts and aggregators exceeds the limit.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMaxActiveResourcesExceededException for service response error code
	// "MaxActiveResourcesExceededException".
	//
	// You have reached the limit (100,000) of active custom resource types in your
	// account. Delete unused resources using DeleteResourceConfig.
	ErrCodeMaxActiveResourcesExceededException = "MaxActiveResourcesExceededException"

	// ErrCodeMaxNumberOfConfigRulesExceededException for service response error code
	// "MaxNumberOfConfigRulesExceededException".
	//
	// Failed to add the AWS Config rule because the account already contains the
	// maximum number of 150 rules. Consider deleting any deactivated rules before
	// you add new rules.
	ErrCodeMaxNumberOfConfigRulesExceededException = "MaxNumberOfConfigRulesExceededException"

	// ErrCodeMaxNumberOfConfigurationRecordersExceededException for service response error code
	// "MaxNumberOfConfigurationRecordersExceededException".
	//
	// You have reached the limit of the number of recorders you can create.
	ErrCodeMaxNumberOfConfigurationRecordersExceededException = "MaxNumberOfConfigurationRecordersExceededException"

	// ErrCodeMaxNumberOfConformancePacksExceededException for service response error code
	// "MaxNumberOfConformancePacksExceededException".
	//
	// You have reached the limit (6) of the number of conformance packs in an account
	// (6 conformance pack with 25 AWS Config rules per pack).
	ErrCodeMaxNumberOfConformancePacksExceededException = "MaxNumberOfConformancePacksExceededException"

	// ErrCodeMaxNumberOfDeliveryChannelsExceededException for service response error code
	// "MaxNumberOfDeliveryChannelsExceededException".
	//
	// You have reached the limit of the number of delivery channels you can create.
	ErrCodeMaxNumberOfDeliveryChannelsExceededException = "MaxNumberOfDeliveryChannelsExceededException"

	// ErrCodeMaxNumberOfOrganizationConfigRulesExceededException for service response error code
	// "MaxNumberOfOrganizationConfigRulesExceededException".
	//
	// You have reached the limit of the number of organization config rules you
	// can create.
	ErrCodeMaxNumberOfOrganizationConfigRulesExceededException = "MaxNumberOfOrganizationConfigRulesExceededException"

	// ErrCodeMaxNumberOfOrganizationConformancePacksExceededException for service response error code
	// "MaxNumberOfOrganizationConformancePacksExceededException".
	//
	// You have reached the limit (6) of the number of organization conformance
	// packs in an account (6 conformance pack with 25 AWS Config rules per pack
	// per account).
	ErrCodeMaxNumberOfOrganizationConformancePacksExceededException = "MaxNumberOfOrganizationConformancePacksExceededException"

	// ErrCodeMaxNumberOfRetentionConfigurationsExceededException for service response error code
	// "MaxNumberOfRetentionConfigurationsExceededException".
	//
	// Failed to add the retention configuration because a retention configuration
	// with that name already exists.
	ErrCodeMaxNumberOfRetentionConfigurationsExceededException = "MaxNumberOfRetentionConfigurationsExceededException"

	// ErrCodeNoAvailableConfigurationRecorderException for service response error code
	// "NoAvailableConfigurationRecorderException".
	//
	// There are no configuration recorders available to provide the role needed
	// to describe your resources. Create a configuration recorder.
	ErrCodeNoAvailableConfigurationRecorderException = "NoAvailableConfigurationRecorderException"

	// ErrCodeNoAvailableDeliveryChannelException for service response error code
	// "NoAvailableDeliveryChannelException".
	//
	// There is no delivery channel available to record configurations.
	ErrCodeNoAvailableDeliveryChannelException = "NoAvailableDeliveryChannelException"

	// ErrCodeNoAvailableOrganizationException for service response error code
	// "NoAvailableOrganizationException".
	//
	// Organization is no longer available.
	ErrCodeNoAvailableOrganizationException = "NoAvailableOrganizationException"

	// ErrCodeNoRunningConfigurationRecorderException for service response error code
	// "NoRunningConfigurationRecorderException".
	//
	// There is no configuration recorder running.
	ErrCodeNoRunningConfigurationRecorderException = "NoRunningConfigurationRecorderException"

	// ErrCodeNoSuchBucketException for service response error code
	// "NoSuchBucketException".
	//
	// The specified Amazon S3 bucket does not exist.
	ErrCodeNoSuchBucketException = "NoSuchBucketException"

	// ErrCodeNoSuchConfigRuleException for service response error code
	// "NoSuchConfigRuleException".
	//
	// One or more AWS Config rules in the request are invalid. Verify that the
	// rule names are correct and try again.
	ErrCodeNoSuchConfigRuleException = "NoSuchConfigRuleException"

	// ErrCodeNoSuchConfigRuleInConformancePackException for service response error code
	// "NoSuchConfigRuleInConformancePackException".
	//
	// AWS Config rule that you passed in the filter does not exist.
	ErrCodeNoSuchConfigRuleInConformancePackException = "NoSuchConfigRuleInConformancePackException"

	// ErrCodeNoSuchConfigurationAggregatorException for service response error code
	// "NoSuchConfigurationAggregatorException".
	//
	// You have specified a configuration aggregator that does not exist.
	ErrCodeNoSuchConfigurationAggregatorException = "NoSuchConfigurationAggregatorException"

	// ErrCodeNoSuchConfigurationRecorderException for service response error code
	// "NoSuchConfigurationRecorderException".
	//
	// You have specified a configuration recorder that does not exist.
	ErrCodeNoSuchConfigurationRecorderException = "NoSuchConfigurationRecorderException"

	// ErrCodeNoSuchConformancePackException for service response error code
	// "NoSuchConformancePackException".
	//
	// You specified one or more conformance packs that do not exist.
	ErrCodeNoSuchConformancePackException = "NoSuchConformancePackException"

	// ErrCodeNoSuchDeliveryChannelException for service response error code
	// "NoSuchDeliveryChannelException".
	//
	// You have specified a delivery channel that does not exist.
	ErrCodeNoSuchDeliveryChannelException = "NoSuchDeliveryChannelException"

	// ErrCodeNoSuchOrganizationConfigRuleException for service response error code
	// "NoSuchOrganizationConfigRuleException".
	//
	// You specified one or more organization config rules that do not exist.
	ErrCodeNoSuchOrganizationConfigRuleException = "NoSuchOrganizationConfigRuleException"

	// ErrCodeNoSuchOrganizationConformancePackException for service response error code
	// "NoSuchOrganizationConformancePackException".
	//
	// AWS Config organization conformance pack that you passed in the filter does
	// not exist.
	//
	// For DeleteOrganizationConformancePack, you tried to delete an organization
	// conformance pack that does not exist.
	ErrCodeNoSuchOrganizationConformancePackException = "NoSuchOrganizationConformancePackException"

	// ErrCodeNoSuchRemediationConfigurationException for service response error code
	// "NoSuchRemediationConfigurationException".
	//
	// You specified an AWS Config rule without a remediation configuration.
	ErrCodeNoSuchRemediationConfigurationException = "NoSuchRemediationConfigurationException"

	// ErrCodeNoSuchRemediationExceptionException for service response error code
	// "NoSuchRemediationExceptionException".
	//
	// You tried to delete a remediation exception that does not exist.
	ErrCodeNoSuchRemediationExceptionException = "NoSuchRemediationExceptionException"

	// ErrCodeNoSuchRetentionConfigurationException for service response error code
	// "NoSuchRetentionConfigurationException".
	//
	// You have specified a retention configuration that does not exist.
	ErrCodeNoSuchRetentionConfigurationException = "NoSuchRetentionConfigurationException"

	// ErrCodeOrganizationAccessDeniedException for service response error code
	// "OrganizationAccessDeniedException".
	//
	// For PutConfigAggregator API, no permission to call EnableAWSServiceAccess
	// API.
	//
	// For all OrganizationConfigRule and OrganizationConformancePack APIs, AWS
	// Config throws an exception if APIs are called from member accounts. All APIs
	// must be called from organization master account.
	ErrCodeOrganizationAccessDeniedException = "OrganizationAccessDeniedException"

	// ErrCodeOrganizationAllFeaturesNotEnabledException for service response error code
	// "OrganizationAllFeaturesNotEnabledException".
	//
	// AWS Config resource cannot be created because your organization does not
	// have all features enabled.
	ErrCodeOrganizationAllFeaturesNotEnabledException = "OrganizationAllFeaturesNotEnabledException"

	// ErrCodeOrganizationConformancePackTemplateValidationException for service response error code
	// "OrganizationConformancePackTemplateValidationException".
	//
	// You have specified a template that is not valid or supported.
	ErrCodeOrganizationConformancePackTemplateValidationException = "OrganizationConformancePackTemplateValidationException"

	// ErrCodeOversizedConfigurationItemException for service response error code
	// "OversizedConfigurationItemException".
	//
	// The configuration item size is outside the allowable range.
	ErrCodeOversizedConfigurationItemException = "OversizedConfigurationItemException"

	// ErrCodeRemediationInProgressException for service response error code
	// "RemediationInProgressException".
	//
	// Remediation action is in progress. You can either cancel execution in AWS
	// Systems Manager or wait and try again later.
	ErrCodeRemediationInProgressException = "RemediationInProgressException"

	// ErrCodeResourceConcurrentModificationException for service response error code
	// "ResourceConcurrentModificationException".
	//
	// Two users are trying to modify the same query at the same time. Wait for
	// a moment and try again.
	ErrCodeResourceConcurrentModificationException = "ResourceConcurrentModificationException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// You see this exception in the following cases:
	//
	//    * For DeleteConfigRule, AWS Config is deleting this rule. Try your request
	//    again later.
	//
	//    * For DeleteConfigRule, the rule is deleting your evaluation results.
	//    Try your request again later.
	//
	//    * For DeleteConfigRule, a remediation action is associated with the rule
	//    and AWS Config cannot delete this rule. Delete the remediation action
	//    associated with the rule before deleting the rule and try your request
	//    again later.
	//
	//    * For PutConfigOrganizationRule, organization config rule deletion is
	//    in progress. Try your request again later.
	//
	//    * For DeleteOrganizationConfigRule, organization config rule creation
	//    is in progress. Try your request again later.
	//
	//    * For PutConformancePack and PutOrganizationConformancePack, a conformance
	//    pack creation, update, and deletion is in progress. Try your request again
	//    later.
	//
	//    * For DeleteConformancePack, a conformance pack creation, update, and
	//    deletion is in progress. Try your request again later.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotDiscoveredException for service response error code
	// "ResourceNotDiscoveredException".
	//
	// You have specified a resource that is either unknown or has not been discovered.
	ErrCodeResourceNotDiscoveredException = "ResourceNotDiscoveredException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// You have specified a resource that does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// You have reached the limit of the number of tags you can use. You have more
	// than 50 tags.
	ErrCodeTooManyTagsException = "TooManyTagsException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The requested action is not valid.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ConformancePackTemplateValidationException":               newErrorConformancePackTemplateValidationException,
	"InsufficientDeliveryPolicyException":                      newErrorInsufficientDeliveryPolicyException,
	"InsufficientPermissionsException":                         newErrorInsufficientPermissionsException,
	"InvalidConfigurationRecorderNameException":                newErrorInvalidConfigurationRecorderNameException,
	"InvalidDeliveryChannelNameException":                      newErrorInvalidDeliveryChannelNameException,
	"InvalidExpressionException":                               newErrorInvalidExpressionException,
	"InvalidLimitException":                                    newErrorInvalidLimitException,
	"InvalidNextTokenException":                                newErrorInvalidNextTokenException,
	"InvalidParameterValueException":                           newErrorInvalidParameterValueException,
	"InvalidRecordingGroupException":                           newErrorInvalidRecordingGroupException,
	"InvalidResultTokenException":                              newErrorInvalidResultTokenException,
	"InvalidRoleException":                                     newErrorInvalidRoleException,
	"InvalidS3KeyPrefixException":                              newErrorInvalidS3KeyPrefixException,
	"InvalidSNSTopicARNException":                              newErrorInvalidSNSTopicARNException,
	"InvalidTimeRangeException":                                newErrorInvalidTimeRangeException,
	"LastDeliveryChannelDeleteFailedException":                 newErrorLastDeliveryChannelDeleteFailedException,
	"LimitExceededException":                                   newErrorLimitExceededException,
	"MaxActiveResourcesExceededException":                      newErrorMaxActiveResourcesExceededException,
	"MaxNumberOfConfigRulesExceededException":                  newErrorMaxNumberOfConfigRulesExceededException,
	"MaxNumberOfConfigurationRecordersExceededException":       newErrorMaxNumberOfConfigurationRecordersExceededException,
	"MaxNumberOfConformancePacksExceededException":             newErrorMaxNumberOfConformancePacksExceededException,
	"MaxNumberOfDeliveryChannelsExceededException":             newErrorMaxNumberOfDeliveryChannelsExceededException,
	"MaxNumberOfOrganizationConfigRulesExceededException":      newErrorMaxNumberOfOrganizationConfigRulesExceededException,
	"MaxNumberOfOrganizationConformancePacksExceededException": newErrorMaxNumberOfOrganizationConformancePacksExceededException,
	"MaxNumberOfRetentionConfigurationsExceededException":      newErrorMaxNumberOfRetentionConfigurationsExceededException,
	"NoAvailableConfigurationRecorderException":                newErrorNoAvailableConfigurationRecorderException,
	"NoAvailableDeliveryChannelException":                      newErrorNoAvailableDeliveryChannelException,
	"NoAvailableOrganizationException":                         newErrorNoAvailableOrganizationException,
	"NoRunningConfigurationRecorderException":                  newErrorNoRunningConfigurationRecorderException,
	"NoSuchBucketException":                                    newErrorNoSuchBucketException,
	"NoSuchConfigRuleException":                                newErrorNoSuchConfigRuleException,
	"NoSuchConfigRuleInConformancePackException":               newErrorNoSuchConfigRuleInConformancePackException,
	"NoSuchConfigurationAggregatorException":                   newErrorNoSuchConfigurationAggregatorException,
	"NoSuchConfigurationRecorderException":                     newErrorNoSuchConfigurationRecorderException,
	"NoSuchConformancePackException":                           newErrorNoSuchConformancePackException,
	"NoSuchDeliveryChannelException":                           newErrorNoSuchDeliveryChannelException,
	"NoSuchOrganizationConfigRuleException":                    newErrorNoSuchOrganizationConfigRuleException,
	"NoSuchOrganizationConformancePackException":               newErrorNoSuchOrganizationConformancePackException,
	"NoSuchRemediationConfigurationException":                  newErrorNoSuchRemediationConfigurationException,
	"NoSuchRemediationExceptionException":                      newErrorNoSuchRemediationExceptionException,
	"NoSuchRetentionConfigurationException":                    newErrorNoSuchRetentionConfigurationException,
	"OrganizationAccessDeniedException":                        newErrorOrganizationAccessDeniedException,
	"OrganizationAllFeaturesNotEnabledException":               newErrorOrganizationAllFeaturesNotEnabledException,
	"OrganizationConformancePackTemplateValidationException":   newErrorOrganizationConformancePackTemplateValidationException,
	"OversizedConfigurationItemException":                      newErrorOversizedConfigurationItemException,
	"RemediationInProgressException":                           newErrorRemediationInProgressException,
	"ResourceConcurrentModificationException":                  newErrorResourceConcurrentModificationException,
	"ResourceInUseException":                                   newErrorResourceInUseException,
	"ResourceNotDiscoveredException":                           newErrorResourceNotDiscoveredException,
	"ResourceNotFoundException":                                newErrorResourceNotFoundException,
	"TooManyTagsException":                                     newErrorTooManyTagsException,
	"ValidationException":                                      newErrorValidationException,
}
