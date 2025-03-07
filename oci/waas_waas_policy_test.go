// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v52/common"
	oci_waas "github.com/oracle/oci-go-sdk/v52/waas"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	waasPolicyDomainSuffix = ".oracle.com"

	waasPolicyDomainName = RandomStringOrHttpReplayValue(4, strings.ToLower(charsetWithoutDigits), "snww")

	waasPolicyDomain = waasPolicyDomainName + waasPolicyDomainSuffix

	WaasPolicyRequiredOnlyResource = WaasPolicyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_waas_waas_policy", "test_waas_policy", Required, Create, waasPolicyRepresentation)

	WaasPolicyResourceConfig = WaasPolicyResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_waas_waas_policy", "test_waas_policy", Optional, Update, waasPolicyRepresentation)

	waasPolicySingularDataSourceRepresentation = map[string]interface{}{
		"waas_policy_id": Representation{RepType: Required, Create: `${oci_waas_waas_policy.test_waas_policy.id}`},
	}

	waasPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_names":                         Representation{RepType: Optional, Create: []string{`displayName2`}},
		"ids":                                   Representation{RepType: Optional, Create: []string{`${oci_waas_waas_policy.test_waas_policy.id}`}},
		"states":                                Representation{RepType: Optional, Create: []string{`ACTIVE`}},
		"time_created_greater_than_or_equal_to": Representation{RepType: Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                Representation{RepType: Optional, Create: `2038-01-01T00:00:00.000Z`},
		"filter":                                RepresentationGroup{Required, waasPolicyDataSourceFilterRepresentation}}
	waasPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_waas_waas_policy.test_waas_policy.id}`}},
	}

	waasPolicyRepresentation = map[string]interface{}{
		"compartment_id":     Representation{RepType: Required, Create: `${var.compartment_id}`},
		"domain":             Representation{RepType: Required, Create: waasPolicyDomain},
		"additional_domains": Representation{RepType: Optional, Create: []string{waasPolicyDomainName + "3" + waasPolicyDomainSuffix, waasPolicyDomainName + "4" + waasPolicyDomainSuffix}, Update: []string{waasPolicyDomainName + "31" + waasPolicyDomainSuffix, waasPolicyDomainName + "41" + waasPolicyDomainSuffix}},
		"defined_tags":       Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":       Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":      Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"origin_groups":      []RepresentationGroup{{Optional, waasOriginGroupsRepresentationMap1}, {Optional, waasOriginGroupsRepresentationMap2}},
		"origins":            []RepresentationGroup{{Optional, waasOriginRepresentationMap1}, {Optional, waasOriginRepresentationMap2}},
		"policy_config":      RepresentationGroup{Optional, waasPolicyPolicyConfigRepresentation},
		"waf_config":         RepresentationGroup{Optional, waasPolicyWafConfigRepresentation},
		"timeouts":           RepresentationGroup{Required, waasPolicyTimeoutsRepresentation},
	}
	waasPolicyTimeoutsRepresentation = map[string]interface{}{
		"create": Representation{RepType: Required, Create: `120m`},
		"update": Representation{RepType: Required, Create: `120m`},
		"delete": Representation{RepType: Required, Create: `120m`},
	}
	waasCustomHeaderRepresentation1 = map[string]interface{}{
		"name":  Representation{RepType: Required, Create: "name1"},
		"value": Representation{RepType: Required, Create: "value1"},
	}
	waasCustomHeaderRepresentation2 = map[string]interface{}{
		"name":  Representation{RepType: Required, Create: "name2"},
		"value": Representation{RepType: Required, Create: "value2"},
	}
	waasOriginGroupRepresentation1 = map[string]interface{}{
		"origin": Representation{RepType: Required, Create: "primary", Update: "primary2"},
		"weight": Representation{RepType: Required, Create: `1`},
	}
	waasOriginGroupRepresentation2 = map[string]interface{}{
		"origin": Representation{RepType: Required, Create: "secondary", Update: "secondary2"},
		"weight": Representation{RepType: Required, Create: `2`},
	}
	waasOriginRepresentationMap1 = map[string]interface{}{
		"label":          Representation{RepType: Required, Create: "primary", Update: "primary2"},
		"uri":            Representation{RepType: Required, Create: "192.168.0.1", Update: "192.168.0.11"},
		"http_port":      Representation{RepType: Required, Create: 80, Update: 8081},
		"https_port":     Representation{RepType: Required, Create: 443, Update: 8444},
		"custom_headers": []RepresentationGroup{{Optional, waasCustomHeaderRepresentation1}, {Optional, waasCustomHeaderRepresentation2}},
	}
	waasOriginGroupsRepresentationMap1 = map[string]interface{}{
		"label":        Representation{RepType: Required, Create: "originGroups1", Update: "originGroups11"},
		"origin_group": []RepresentationGroup{{Optional, waasOriginGroupRepresentation1}, {Optional, waasOriginGroupRepresentation2}},
	}
	waasOriginGroupsRepresentationMap2 = map[string]interface{}{
		"label":        Representation{RepType: Required, Create: "originGroups2", Update: "originGroups22"},
		"origin_group": []RepresentationGroup{{Optional, waasOriginGroupRepresentation1}, {Optional, waasOriginGroupRepresentation2}},
	}
	waasOriginRepresentationMap2 = map[string]interface{}{
		"label":          Representation{RepType: Required, Create: "secondary", Update: "secondary2"},
		"uri":            Representation{RepType: Required, Create: "192.168.0.2", Update: "192.168.0.20"},
		"http_port":      Representation{RepType: Required, Create: 8080, Update: 8082},
		"https_port":     Representation{RepType: Required, Create: 8443, Update: 8445},
		"custom_headers": []RepresentationGroup{{Optional, waasCustomHeaderRepresentation1}, {Optional, waasCustomHeaderRepresentation2}},
	}
	waasPolicyPolicyConfigRepresentation = map[string]interface{}{
		"certificate_id":                Representation{RepType: Optional, Create: `${oci_waas_certificate.test_certificate.id}`},
		"cipher_group":                  Representation{RepType: Optional, Create: `DEFAULT`, Update: `DEFAULT`},
		"client_address_header":         Representation{RepType: Optional, Create: ``, Update: `X_FORWARDED_FOR`},
		"health_checks":                 RepresentationGroup{Optional, waasPolicyPolicyConfigHealthChecksRepresentation},
		"is_behind_cdn":                 Representation{RepType: Optional, Create: `false`, Update: `true`},
		"is_cache_control_respected":    Representation{RepType: Optional, Create: `false`, Update: `true`},
		"is_https_enabled":              Representation{RepType: Optional, Create: `false`, Update: `true`},
		"is_https_forced":               Representation{RepType: Optional, Create: `false`, Update: `true`},
		"is_origin_compression_enabled": Representation{RepType: Optional, Create: `false`, Update: `true`},
		"is_response_buffering_enabled": Representation{RepType: Optional, Create: `false`, Update: `true`},
		"is_sni_enabled":                Representation{RepType: Optional, Create: `false`, Update: `true`},
		"load_balancing_method":         RepresentationGroup{Optional, waasPolicyPolicyConfigLoadBalancingMethodRepresentation},
		"tls_protocols":                 Representation{RepType: Optional, Create: []string{`TLS_V1_2`}, Update: []string{`TLS_V1_3`}},
		"websocket_path_prefixes":       Representation{RepType: Optional, Create: []string{`/url1`}, Update: []string{`/url2`}},
	}
	waasPolicyWafConfigRepresentation = map[string]interface{}{
		"access_rules":          RepresentationGroup{Optional, waasPolicyWafConfigAccessRulesRepresentation},
		"address_rate_limiting": RepresentationGroup{Optional, waasPolicyWafConfigAddressRateLimitingRepresentation},
		"caching_rules":         RepresentationGroup{Optional, waasPolicyWafConfigCachingRulesRepresentation},
		"captchas":              RepresentationGroup{Optional, waasPolicyWafConfigCaptchasRepresentation},
		//@Codegen: awaiting resolution for the known issue of deletion wait time for linked customProtectionRule to a policy
		//"custom_protection_rules":      RepresentationGroup{Optional, waasPolicyWafConfigCustomProtectionRulesRepresentation},
		"device_fingerprint_challenge": RepresentationGroup{Optional, waasPolicyWafConfigDeviceFingerprintChallengeRepresentation},
		"human_interaction_challenge":  RepresentationGroup{Optional, waasPolicyWafConfigHumanInteractionChallengeRepresentation},
		"js_challenge":                 RepresentationGroup{Optional, waasPolicyWafConfigJsChallengeRepresentation},
		"origin":                       Representation{RepType: Optional, Create: `primary`, Update: `primary2`},
		"origin_groups":                Representation{RepType: Optional, Create: []string{`originGroups1`}, Update: []string{`originGroups11`}},
		"protection_settings":          RepresentationGroup{Optional, waasPolicyWafConfigProtectionSettingsRepresentation},
		"whitelists":                   RepresentationGroup{Optional, waasPolicyWafConfigWhitelistsRepresentation},
	}
	waasPolicyOriginsCustomHeadersRepresentation = map[string]interface{}{
		"name":  Representation{RepType: Required, Create: `name`, Update: `name2`},
		"value": Representation{RepType: Required, Create: `value`, Update: `value2`},
	}
	waasPolicyPolicyConfigHealthChecksRepresentation = map[string]interface{}{
		"expected_response_code_group":   Representation{RepType: Optional, Create: []string{`2XX`}, Update: []string{`3XX`}},
		"expected_response_text":         Representation{RepType: Optional, Create: `expectedResponseText`, Update: `expectedResponseText2`},
		"headers":                        Representation{RepType: Optional, Create: map[string]string{"Host": "oracle.com", "User-Agent": "Oracle-TerraformProvider"}},
		"healthy_threshold":              Representation{RepType: Optional, Create: `10`, Update: `11`},
		"interval_in_seconds":            Representation{RepType: Optional, Create: `10`, Update: `11`},
		"is_enabled":                     Representation{RepType: Optional, Create: `false`, Update: `true`},
		"is_response_text_check_enabled": Representation{RepType: Optional, Create: `false`, Update: `true`},
		"method":                         Representation{RepType: Optional, Create: `GET`, Update: `POST`},
		"path":                           Representation{RepType: Optional, Create: `/`},
		"timeout_in_seconds":             Representation{RepType: Optional, Create: `10`, Update: `11`},
		"unhealthy_threshold":            Representation{RepType: Optional, Create: `10`, Update: `11`},
	}
	waasPolicyPolicyConfigLoadBalancingMethodRepresentation = map[string]interface{}{
		"method":                     Representation{RepType: Required, Create: `STICKY_COOKIE`},
		"domain":                     Representation{RepType: Optional, Create: `example.com`, Update: `example2.com`},
		"expiration_time_in_seconds": Representation{RepType: Optional, Create: `10`, Update: `11`},
		"name":                       Representation{RepType: Optional, Create: `name`, Update: `name2`},
	}
	waasPolicyWafConfigAccessRulesRepresentation = map[string]interface{}{
		"action":                       Representation{RepType: Required, Create: `ALLOW`, Update: `DETECT`},
		"criteria":                     RepresentationGroup{Required, waasPolicyWafConfigAccessRulesCriteriaRepresentation},
		"name":                         Representation{RepType: Required, Create: `name`, Update: `name2`},
		"block_action":                 Representation{RepType: Optional, Create: `SET_RESPONSE_CODE`, Update: `SHOW_ERROR_PAGE`},
		"block_error_page_code":        Representation{RepType: Optional, Create: `403`, Update: `401`},
		"block_error_page_description": Representation{RepType: Optional, Create: `blockErrorPageDescription`, Update: `blockErrorPageDescription2`},
		"block_error_page_message":     Representation{RepType: Optional, Create: `blockErrorPageMessage`, Update: `blockErrorPageMessage2`},
		"block_response_code":          Representation{RepType: Optional, Create: `403`, Update: `401`},
		"bypass_challenges":            Representation{RepType: Optional, Create: []string{`JS_CHALLENGE`}, Update: []string{`HUMAN_INTERACTION_CHALLENGE`}},
		"captcha_footer":               Representation{RepType: Optional, Create: `captchaFooter`, Update: `captchaFooter2`},
		"captcha_header":               Representation{RepType: Optional, Create: `captchaHeader`, Update: `captchaHeader2`},
		"captcha_submit_label":         Representation{RepType: Optional, Create: `captchaSubmitLabel`, Update: `captchaSubmitLabel2`},
		"captcha_title":                Representation{RepType: Optional, Create: `captchaTitle`, Update: `captchaTitle2`},
		"redirect_response_code":       Representation{RepType: Optional, Create: `FOUND`, Update: `MOVED_PERMANENTLY`},
		"redirect_url":                 Representation{RepType: Optional, Create: `http://0.0.0.0:80`, Update: `http://0.0.0.0:81`},
		"response_header_manipulation": RepresentationGroup{Optional, waasPolicyWafConfigAccessRulesResponseHeaderManipulationRepresentation},
	}
	waasPolicyWafConfigAddressRateLimitingRepresentation = map[string]interface{}{
		"is_enabled":                    Representation{RepType: Required, Create: `false`, Update: `true`},
		"allowed_rate_per_address":      Representation{RepType: Optional, Create: `10`, Update: `11`},
		"block_response_code":           Representation{RepType: Optional, Create: `403`, Update: `401`},
		"max_delayed_count_per_address": Representation{RepType: Optional, Create: `10`, Update: `11`},
	}
	waasPolicyWafConfigCachingRulesRepresentation = map[string]interface{}{
		"action":                    Representation{RepType: Required, Create: `CACHE`, Update: `CACHE`},
		"criteria":                  RepresentationGroup{Required, waasPolicyWafConfigCachingRulesCriteriaRepresentation},
		"name":                      Representation{RepType: Required, Create: `name`, Update: `name2`},
		"caching_duration":          Representation{RepType: Optional, Create: `PT1S`, Update: `PT2S`},
		"client_caching_duration":   Representation{RepType: Optional, Create: `PT1S`, Update: `PT2S`},
		"is_client_caching_enabled": Representation{RepType: Optional, Create: `false`, Update: `true`},
		"key":                       Representation{RepType: Optional, Create: `key`, Update: `key2`},
	}
	waasPolicyWafConfigCaptchasRepresentation = map[string]interface{}{
		"failure_message":               Representation{RepType: Required, Create: `failureMessage`, Update: `failureMessage2`},
		"session_expiration_in_seconds": Representation{RepType: Required, Create: `10`, Update: `11`},
		"submit_label":                  Representation{RepType: Required, Create: `submitLabel`, Update: `submitLabel2`},
		"title":                         Representation{RepType: Required, Create: `title`, Update: `title2`},
		"url":                           Representation{RepType: Required, Create: `url`, Update: `url2`},
		"footer_text":                   Representation{RepType: Optional, Create: `footerText`, Update: `footerText2`},
		"header_text":                   Representation{RepType: Optional, Create: `headerText`, Update: `headerText2`},
	}
	waasPolicyWafConfigCustomProtectionRulesRepresentation = map[string]interface{}{
		"action":     Representation{RepType: Optional, Create: `DETECT`, Update: `BLOCK`},
		"exclusions": RepresentationGroup{Optional, waasPolicyWafConfigCustomProtectionRulesExclusionsRepresentation},
		"id":         Representation{RepType: Optional, Create: `${oci_waas_custom_protection_rule.test_custom_protection_rule.id}`, Update: `${oci_waas_custom_protection_rule.test_custom_protection_rule2.id}`},
	}
	waasPolicyWafConfigDeviceFingerprintChallengeRepresentation = map[string]interface{}{
		"is_enabled":                   Representation{RepType: Required, Create: `false`, Update: `true`},
		"action":                       Representation{RepType: Optional, Create: `DETECT`, Update: `BLOCK`},
		"action_expiration_in_seconds": Representation{RepType: Optional, Create: `10`, Update: `11`},
		"challenge_settings":           RepresentationGroup{Optional, waasPolicyWafConfigDeviceFingerprintChallengeChallengeSettingsRepresentation},
		"failure_threshold":            Representation{RepType: Optional, Create: `10`, Update: `11`},
		"failure_threshold_expiration_in_seconds": Representation{RepType: Optional, Create: `10`, Update: `11`},
		"max_address_count":                       Representation{RepType: Optional, Create: `10`, Update: `11`},
		"max_address_count_expiration_in_seconds": Representation{RepType: Optional, Create: `10`, Update: `11`},
	}
	waasPolicyWafConfigHumanInteractionChallengeRepresentation = map[string]interface{}{
		"is_enabled":                   Representation{RepType: Required, Create: `false`, Update: `true`},
		"action":                       Representation{RepType: Optional, Create: `DETECT`, Update: `BLOCK`},
		"action_expiration_in_seconds": Representation{RepType: Optional, Create: `10`, Update: `11`},
		"challenge_settings":           RepresentationGroup{Optional, waasPolicyWafConfigHumanInteractionChallengeChallengeSettingsRepresentation},
		"failure_threshold":            Representation{RepType: Optional, Create: `10`, Update: `11`},
		"failure_threshold_expiration_in_seconds": Representation{RepType: Optional, Create: `10`, Update: `11`},
		"interaction_threshold":                   Representation{RepType: Optional, Create: `10`, Update: `11`},
		"is_nat_enabled":                          Representation{RepType: Optional, Create: `false`, Update: `true`},
		"recording_period_in_seconds":             Representation{RepType: Optional, Create: `10`, Update: `11`},
		"set_http_header":                         RepresentationGroup{Optional, waasPolicyWafConfigHumanInteractionChallengeSetHttpHeaderRepresentation},
	}
	waasPolicyWafConfigJsChallengeRepresentation = map[string]interface{}{
		"is_enabled":                   Representation{RepType: Required, Create: `false`, Update: `true`},
		"action":                       Representation{RepType: Optional, Create: `DETECT`, Update: `BLOCK`},
		"action_expiration_in_seconds": Representation{RepType: Optional, Create: `10`, Update: `11`},
		"are_redirects_challenged":     Representation{RepType: Optional, Create: `false`, Update: `true`},
		"challenge_settings":           RepresentationGroup{Optional, waasPolicyWafConfigJsChallengeChallengeSettingsRepresentation},
		"criteria":                     RepresentationGroup{Optional, waasPolicyWafConfigJsChallengeCriteriaRepresentation},
		"failure_threshold":            Representation{RepType: Optional, Create: `10`, Update: `11`},
		"is_nat_enabled":               Representation{RepType: Optional, Create: `false`, Update: `true`},
		"set_http_header":              RepresentationGroup{Optional, waasPolicyWafConfigJsChallengeSetHttpHeaderRepresentation},
	}
	waasPolicyWafConfigProtectionSettingsRepresentation = map[string]interface{}{
		"allowed_http_methods":               Representation{RepType: Optional, Create: []string{`OPTIONS`}, Update: []string{`HEAD`}},
		"block_action":                       Representation{RepType: Optional, Create: `SET_RESPONSE_CODE`, Update: `SHOW_ERROR_PAGE`},
		"block_error_page_code":              Representation{RepType: Optional, Create: `403`, Update: `401`},
		"block_error_page_description":       Representation{RepType: Optional, Create: `blockErrorPageDescription`, Update: `blockErrorPageDescription2`},
		"block_error_page_message":           Representation{RepType: Optional, Create: `blockErrorPageMessage`, Update: `blockErrorPageMessage2`},
		"block_response_code":                Representation{RepType: Optional, Create: `403`, Update: `401`},
		"is_response_inspected":              Representation{RepType: Optional, Create: `false`},
		"max_argument_count":                 Representation{RepType: Optional, Create: `10`, Update: `11`},
		"max_name_length_per_argument":       Representation{RepType: Optional, Create: `10`, Update: `11`},
		"max_response_size_in_ki_b":          Representation{RepType: Optional, Create: `10`, Update: `11`},
		"max_total_name_length_of_arguments": Representation{RepType: Optional, Create: `10`, Update: `11`},
		"media_types":                        Representation{RepType: Optional, Create: []string{`application/plain`}, Update: []string{`application/json`}},
		"recommendations_period_in_days":     Representation{RepType: Optional, Create: `10`, Update: `11`},
	}
	waasPolicyWafConfigWhitelistsRepresentation = map[string]interface{}{
		"name":      Representation{RepType: Required, Create: `name`, Update: `name2`},
		"addresses": Representation{RepType: Optional, Create: []string{`192.168.127.127`}, Update: []string{`192.168.127.128`}},
	}
	waasPolicyWafConfigAccessRulesCriteriaRepresentation = map[string]interface{}{
		"condition":         Representation{RepType: Required, Create: `URL_IS`, Update: `URL_IS_NOT`},
		"value":             Representation{RepType: Required, Create: `/public`, Update: `/secret`},
		"is_case_sensitive": Representation{RepType: Optional, Create: `false`, Update: `true`},
	}
	waasPolicyWafConfigAccessRulesResponseHeaderManipulationRepresentation = map[string]interface{}{
		"action": Representation{RepType: Required, Create: `EXTEND_HTTP_RESPONSE_HEADER`, Update: `ADD_HTTP_RESPONSE_HEADER`},
		"header": Representation{RepType: Required, Create: `header`, Update: `header2`},
		"value":  Representation{RepType: Optional, Create: `value`, Update: `value2`},
	}
	waasPolicyWafConfigCachingRulesCriteriaRepresentation = map[string]interface{}{
		"condition": Representation{RepType: Required, Create: `URL_IS`, Update: `URL_STARTS_WITH`},
		"value":     Representation{RepType: Required, Create: `/public`, Update: `/publ`},
	}
	waasPolicyWafConfigCustomProtectionRulesExclusionsRepresentation = map[string]interface{}{
		"exclusions": Representation{RepType: Optional, Create: []string{`example.com`}, Update: []string{`example2.com`}},
		"target":     Representation{RepType: Optional, Create: `REQUEST_COOKIES`, Update: `target2`},
	}
	waasPolicyWafConfigDeviceFingerprintChallengeChallengeSettingsRepresentation = map[string]interface{}{
		"block_action":                 Representation{RepType: Optional, Create: `SET_RESPONSE_CODE`, Update: `SHOW_ERROR_PAGE`},
		"block_error_page_code":        Representation{RepType: Optional, Create: `403`, Update: `401`},
		"block_error_page_description": Representation{RepType: Optional, Create: `blockErrorPageDescription`, Update: `blockErrorPageDescription2`},
		"block_error_page_message":     Representation{RepType: Optional, Create: `blockErrorPageMessage`, Update: `blockErrorPageMessage2`},
		"block_response_code":          Representation{RepType: Optional, Create: `403`, Update: `401`},
		"captcha_footer":               Representation{RepType: Optional, Create: `captchaFooter`, Update: `captchaFooter2`},
		"captcha_header":               Representation{RepType: Optional, Create: `captchaHeader`, Update: `captchaHeader2`},
		"captcha_submit_label":         Representation{RepType: Optional, Create: `captchaSubmitLabel`, Update: `captchaSubmitLabel2`},
		"captcha_title":                Representation{RepType: Optional, Create: `captchaTitle`, Update: `captchaTitle2`},
	}
	waasPolicyWafConfigHumanInteractionChallengeChallengeSettingsRepresentation = map[string]interface{}{
		"block_action":                 Representation{RepType: Optional, Create: `SET_RESPONSE_CODE`, Update: `SHOW_ERROR_PAGE`},
		"block_error_page_code":        Representation{RepType: Optional, Create: `403`, Update: `401`},
		"block_error_page_description": Representation{RepType: Optional, Create: `blockErrorPageDescription`, Update: `blockErrorPageDescription2`},
		"block_error_page_message":     Representation{RepType: Optional, Create: `blockErrorPageMessage`, Update: `blockErrorPageMessage2`},
		"block_response_code":          Representation{RepType: Optional, Create: `403`, Update: `401`},
		"captcha_footer":               Representation{RepType: Optional, Create: `captchaFooter`, Update: `captchaFooter2`},
		"captcha_header":               Representation{RepType: Optional, Create: `captchaHeader`, Update: `captchaHeader2`},
		"captcha_submit_label":         Representation{RepType: Optional, Create: `captchaSubmitLabel`, Update: `captchaSubmitLabel2`},
		"captcha_title":                Representation{RepType: Optional, Create: `captchaTitle`, Update: `captchaTitle2`},
	}
	waasPolicyWafConfigHumanInteractionChallengeSetHttpHeaderRepresentation = map[string]interface{}{
		"name":  Representation{RepType: Required, Create: `name`, Update: `name2`},
		"value": Representation{RepType: Required, Create: `value`, Update: `value2`},
	}
	waasPolicyWafConfigJsChallengeChallengeSettingsRepresentation = map[string]interface{}{
		"block_action":                 Representation{RepType: Optional, Create: `SET_RESPONSE_CODE`, Update: `SHOW_ERROR_PAGE`},
		"block_error_page_code":        Representation{RepType: Optional, Create: `403`, Update: `401`},
		"block_error_page_description": Representation{RepType: Optional, Create: `blockErrorPageDescription`, Update: `blockErrorPageDescription2`},
		"block_error_page_message":     Representation{RepType: Optional, Create: `blockErrorPageMessage`, Update: `blockErrorPageMessage2`},
		"block_response_code":          Representation{RepType: Optional, Create: `403`, Update: `401`},
		"captcha_footer":               Representation{RepType: Optional, Create: `captchaFooter`, Update: `captchaFooter2`},
		"captcha_header":               Representation{RepType: Optional, Create: `captchaHeader`, Update: `captchaHeader2`},
		"captcha_submit_label":         Representation{RepType: Optional, Create: `captchaSubmitLabel`, Update: `captchaSubmitLabel2`},
		"captcha_title":                Representation{RepType: Optional, Create: `captchaTitle`, Update: `captchaTitle2`},
	}
	waasPolicyWafConfigJsChallengeCriteriaRepresentation = map[string]interface{}{
		"condition":         Representation{RepType: Required, Create: `URL_IS`, Update: `URL_STARTS_WITH`},
		"value":             Representation{RepType: Required, Create: `/public`, Update: `/publ`},
		"is_case_sensitive": Representation{RepType: Optional, Create: `false`, Update: `true`},
	}
	waasPolicyWafConfigJsChallengeSetHttpHeaderRepresentation = map[string]interface{}{
		"name":  Representation{RepType: Required, Create: `name`, Update: `name2`},
		"value": Representation{RepType: Required, Create: `value`, Update: `value2`},
	}

	WaasPolicyResourceDependencies = DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_waas_certificate", "test_certificate", Required, Create, waasCertificateRepresentation) +
		caCertificateVariableStr +
		privateKeyVariableStr +
		CustomProtectionRuleRequiredResourceWithoutDependencies
)

// issue-routing-tag: waas/default
func TestWaasWaasPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWaasWaasPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_waas_waas_policy.test_waas_policy"
	datasourceName := "data.oci_waas_waas_policies.test_waas_policies"
	singularDatasourceName := "data.oci_waas_waas_policy.test_waas_policy"

	var resId, resId2 string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+WaasPolicyResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_waas_waas_policy", "test_waas_policy", Optional, Create, waasPolicyRepresentation), "waas", "waasPolicy", t)

	ResourceTest(t, testAccCheckWaasWaasPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + WaasPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waas_waas_policy", "test_waas_policy", Required, Create, waasPolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "domain", waasPolicyDomain),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + WaasPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + WaasPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waas_waas_policy", "test_waas_policy", Optional, Create, waasPolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "additional_domains.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "domain", waasPolicyDomain),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "origin_groups.#", "2"),
				CheckResourceSetContainsElementWithProperties(resourceName, "origin_groups", map[string]string{
					"origin_group.#": "2",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "origins.#", "2"),
				CheckResourceSetContainsElementWithProperties(resourceName, "origins", map[string]string{
					"custom_headers.#": "2",
					"http_port":        "80",
					"https_port":       "443",
					"uri":              "192.168.0.1",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "policy_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "policy_config.0.certificate_id"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.cipher_group", "DEFAULT"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.client_address_header", ""),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.expected_response_code_group.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.expected_response_text", "expectedResponseText"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.headers.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.healthy_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.interval_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.is_response_text_check_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.method", "GET"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.path", "/"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.timeout_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.unhealthy_threshold", "10"),

				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_behind_cdn", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_cache_control_respected", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_https_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_https_forced", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_origin_compression_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_response_buffering_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_sni_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.load_balancing_method.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.load_balancing_method.0.domain", "example.com"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.load_balancing_method.0.expiration_time_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.load_balancing_method.0.method", "STICKY_COOKIE"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.load_balancing_method.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.tls_protocols.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.websocket_path_prefixes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.action", "ALLOW"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.block_action", "SET_RESPONSE_CODE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.block_error_page_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.block_error_page_description", "blockErrorPageDescription"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.block_error_page_message", "blockErrorPageMessage"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.block_response_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.bypass_challenges.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.captcha_footer", "captchaFooter"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.captcha_header", "captchaHeader"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.captcha_submit_label", "captchaSubmitLabel"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.captcha_title", "captchaTitle"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.criteria.0.condition", "URL_IS"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.criteria.0.is_case_sensitive", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.criteria.0.value", "/public"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.redirect_response_code", "FOUND"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.redirect_url", "http://0.0.0.0:80"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.response_header_manipulation.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.response_header_manipulation.0.action", "EXTEND_HTTP_RESPONSE_HEADER"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.response_header_manipulation.0.header", "header"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.response_header_manipulation.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.address_rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.address_rate_limiting.0.allowed_rate_per_address", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.address_rate_limiting.0.block_response_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.address_rate_limiting.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.address_rate_limiting.0.max_delayed_count_per_address", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.action", "CACHE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.caching_duration", "PT1S"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.client_caching_duration", "PT1S"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.criteria.0.condition", "URL_IS"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.criteria.0.value", "/public"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.is_client_caching_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.failure_message", "failureMessage"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.footer_text", "footerText"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.header_text", "headerText"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.session_expiration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.submit_label", "submitLabel"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.title", "title"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.url", "url"),
				//resource.TestCheckResourceAttr(resourceName, "waf_config.0.custom_protection_rules.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "waf_config.0.custom_protection_rules.0.action", "DETECT"),
				//resource.TestCheckResourceAttrSet(resourceName, "waf_config.0.custom_protection_rules.0.id"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.action", "DETECT"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.action_expiration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_action", "SET_RESPONSE_CODE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_error_page_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_error_page_description", "blockErrorPageDescription"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_error_page_message", "blockErrorPageMessage"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_response_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_footer", "captchaFooter"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_header", "captchaHeader"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_submit_label", "captchaSubmitLabel"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_title", "captchaTitle"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.failure_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.failure_threshold_expiration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.max_address_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.max_address_count_expiration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.action", "DETECT"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.action_expiration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_action", "SET_RESPONSE_CODE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_error_page_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_error_page_description", "blockErrorPageDescription"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_error_page_message", "blockErrorPageMessage"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_response_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_footer", "captchaFooter"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_header", "captchaHeader"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_submit_label", "captchaSubmitLabel"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_title", "captchaTitle"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.failure_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.failure_threshold_expiration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.interaction_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.is_nat_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.recording_period_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.set_http_header.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.set_http_header.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.set_http_header.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.action", "DETECT"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.action_expiration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.are_redirects_challenged", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_action", "SET_RESPONSE_CODE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_error_page_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_error_page_description", "blockErrorPageDescription"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_error_page_message", "blockErrorPageMessage"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_response_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_footer", "captchaFooter"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_header", "captchaHeader"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_submit_label", "captchaSubmitLabel"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_title", "captchaTitle"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.criteria.0.condition", "URL_IS"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.criteria.0.is_case_sensitive", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.criteria.0.value", "/public"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.failure_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.is_nat_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.set_http_header.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.set_http_header.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.set_http_header.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.origin", "primary"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.origin_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.allowed_http_methods.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.block_action", "SET_RESPONSE_CODE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.block_error_page_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.block_error_page_description", "blockErrorPageDescription"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.block_error_page_message", "blockErrorPageMessage"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.block_response_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.is_response_inspected", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.max_argument_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.max_name_length_per_argument", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.max_response_size_in_ki_b", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.max_total_name_length_of_arguments", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.media_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.recommendations_period_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.whitelists.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.whitelists.0.addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.whitelists.0.name", "name"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + WaasPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waas_waas_policy", "test_waas_policy", Optional, Create,
					RepresentationCopyWithNewProperties(waasPolicyRepresentation, map[string]interface{}{
						"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "additional_domains.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "domain", waasPolicyDomain),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "origin_groups.#", "2"),
				CheckResourceSetContainsElementWithProperties(resourceName, "origin_groups", map[string]string{
					"origin_group.#": "2",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "origins.#", "2"),
				CheckResourceSetContainsElementWithProperties(resourceName, "origins", map[string]string{
					"custom_headers.#": "2",
					"http_port":        "80",
					"https_port":       "443",
					"uri":              "192.168.0.1",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "policy_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "policy_config.0.certificate_id"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.cipher_group", "DEFAULT"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.expected_response_code_group.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.expected_response_text", "expectedResponseText"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.headers.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.healthy_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.interval_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.is_response_text_check_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.method", "GET"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.path", "/"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.timeout_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.unhealthy_threshold", "10"),
				//resource.TestCheckResourceAttr(resourceName, "policy_config.0.client_address_header", "X_FORWARDED_FOR"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_behind_cdn", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_cache_control_respected", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_https_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_https_forced", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_origin_compression_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_response_buffering_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_sni_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.load_balancing_method.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.load_balancing_method.0.domain", "example.com"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.load_balancing_method.0.expiration_time_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.load_balancing_method.0.method", "STICKY_COOKIE"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.load_balancing_method.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.tls_protocols.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.websocket_path_prefixes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.action", "ALLOW"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.block_action", "SET_RESPONSE_CODE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.block_error_page_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.block_error_page_description", "blockErrorPageDescription"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.block_error_page_message", "blockErrorPageMessage"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.block_response_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.bypass_challenges.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.captcha_footer", "captchaFooter"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.captcha_header", "captchaHeader"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.captcha_submit_label", "captchaSubmitLabel"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.captcha_title", "captchaTitle"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.criteria.0.condition", "URL_IS"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.criteria.0.is_case_sensitive", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.criteria.0.value", "/public"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.redirect_response_code", "FOUND"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.redirect_url", "http://0.0.0.0:80"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.response_header_manipulation.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.response_header_manipulation.0.action", "EXTEND_HTTP_RESPONSE_HEADER"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.response_header_manipulation.0.header", "header"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.response_header_manipulation.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.address_rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.address_rate_limiting.0.allowed_rate_per_address", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.address_rate_limiting.0.block_response_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.address_rate_limiting.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.address_rate_limiting.0.max_delayed_count_per_address", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.action", "CACHE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.caching_duration", "PT1S"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.client_caching_duration", "PT1S"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.criteria.0.condition", "URL_IS"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.criteria.0.value", "/public"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.is_client_caching_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.key", "key"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.failure_message", "failureMessage"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.footer_text", "footerText"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.header_text", "headerText"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.session_expiration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.submit_label", "submitLabel"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.title", "title"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.url", "url"),
				//resource.TestCheckResourceAttr(resourceName, "waf_config.0.custom_protection_rules.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "waf_config.0.custom_protection_rules.0.action", "DETECT"),
				//resource.TestCheckResourceAttrSet(resourceName, "waf_config.0.custom_protection_rules.0.id"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.action", "DETECT"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.action_expiration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_action", "SET_RESPONSE_CODE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_error_page_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_error_page_description", "blockErrorPageDescription"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_error_page_message", "blockErrorPageMessage"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_response_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_footer", "captchaFooter"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_header", "captchaHeader"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_submit_label", "captchaSubmitLabel"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_title", "captchaTitle"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.failure_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.failure_threshold_expiration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.max_address_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.max_address_count_expiration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.action", "DETECT"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.action_expiration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_action", "SET_RESPONSE_CODE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_error_page_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_error_page_description", "blockErrorPageDescription"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_error_page_message", "blockErrorPageMessage"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_response_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_footer", "captchaFooter"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_header", "captchaHeader"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_submit_label", "captchaSubmitLabel"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_title", "captchaTitle"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.failure_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.failure_threshold_expiration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.interaction_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.is_nat_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.recording_period_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.set_http_header.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.set_http_header.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.set_http_header.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.action", "DETECT"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.action_expiration_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.are_redirects_challenged", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_action", "SET_RESPONSE_CODE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_error_page_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_error_page_description", "blockErrorPageDescription"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_error_page_message", "blockErrorPageMessage"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_response_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_footer", "captchaFooter"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_header", "captchaHeader"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_submit_label", "captchaSubmitLabel"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_title", "captchaTitle"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.criteria.0.condition", "URL_IS"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.criteria.0.is_case_sensitive", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.criteria.0.value", "/public"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.failure_threshold", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.is_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.is_nat_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.set_http_header.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.set_http_header.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.set_http_header.0.value", "value"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.origin", "primary"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.origin_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.allowed_http_methods.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.block_action", "SET_RESPONSE_CODE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.block_error_page_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.block_error_page_description", "blockErrorPageDescription"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.block_error_page_message", "blockErrorPageMessage"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.block_response_code", "403"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.is_response_inspected", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.max_argument_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.max_name_length_per_argument", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.max_response_size_in_ki_b", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.max_total_name_length_of_arguments", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.media_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.recommendations_period_in_days", "10"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.whitelists.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.whitelists.0.addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.whitelists.0.name", "name"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + WaasPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waas_waas_policy", "test_waas_policy", Optional, Update, waasPolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "additional_domains.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "domain", waasPolicyDomain),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "origin_groups.#", "2"),
				CheckResourceSetContainsElementWithProperties(resourceName, "origin_groups", map[string]string{
					"origin_group.#": "2",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "origins.#", "2"),
				CheckResourceSetContainsElementWithProperties(resourceName, "origins", map[string]string{
					"custom_headers.#": "2",
					"http_port":        "80",
					"https_port":       "443",
					"uri":              "192.168.0.11",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "policy_config.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "policy_config.0.certificate_id"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.cipher_group", "DEFAULT"),

				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.expected_response_code_group.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.expected_response_text", "expectedResponseText2"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.headers.%", "2"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.healthy_threshold", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.interval_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.is_response_text_check_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.method", "POST"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.path", "/"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.timeout_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.health_checks.0.unhealthy_threshold", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.client_address_header", "X_FORWARDED_FOR"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_behind_cdn", "true"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_cache_control_respected", "true"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_https_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_https_forced", "true"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_origin_compression_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_response_buffering_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.is_sni_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.load_balancing_method.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.load_balancing_method.0.domain", "example2.com"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.load_balancing_method.0.expiration_time_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.load_balancing_method.0.method", "STICKY_COOKIE"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.load_balancing_method.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.tls_protocols.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "policy_config.0.websocket_path_prefixes.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.action", "DETECT"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.block_action", "SHOW_ERROR_PAGE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.block_error_page_code", "401"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.block_error_page_description", "blockErrorPageDescription2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.block_error_page_message", "blockErrorPageMessage2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.block_response_code", "401"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.bypass_challenges.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.captcha_footer", "captchaFooter2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.captcha_header", "captchaHeader2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.captcha_submit_label", "captchaSubmitLabel2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.captcha_title", "captchaTitle2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.criteria.0.condition", "URL_IS_NOT"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.criteria.0.is_case_sensitive", "true"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.criteria.0.value", "/secret"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.redirect_response_code", "MOVED_PERMANENTLY"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.redirect_url", "http://0.0.0.0:81"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.response_header_manipulation.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.response_header_manipulation.0.action", "ADD_HTTP_RESPONSE_HEADER"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.response_header_manipulation.0.header", "header2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.access_rules.0.response_header_manipulation.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.address_rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.address_rate_limiting.0.allowed_rate_per_address", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.address_rate_limiting.0.block_response_code", "401"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.address_rate_limiting.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.address_rate_limiting.0.max_delayed_count_per_address", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.action", "CACHE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.caching_duration", "PT2S"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.client_caching_duration", "PT2S"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.criteria.0.condition", "URL_STARTS_WITH"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.criteria.0.value", "/publ"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.is_client_caching_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.key", "key2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.caching_rules.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.failure_message", "failureMessage2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.footer_text", "footerText2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.header_text", "headerText2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.session_expiration_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.submit_label", "submitLabel2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.title", "title2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.captchas.0.url", "url2"),
				//resource.TestCheckResourceAttr(resourceName, "waf_config.0.custom_protection_rules.#", "1"),
				//resource.TestCheckResourceAttr(resourceName, "waf_config.0.custom_protection_rules.0.action", "BLOCK"),
				//resource.TestCheckResourceAttrSet(resourceName, "waf_config.0.custom_protection_rules.0.id"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.action", "BLOCK"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.action_expiration_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_action", "SHOW_ERROR_PAGE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_error_page_code", "401"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_error_page_description", "blockErrorPageDescription2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_error_page_message", "blockErrorPageMessage2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_response_code", "401"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_footer", "captchaFooter2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_header", "captchaHeader2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_submit_label", "captchaSubmitLabel2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_title", "captchaTitle2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.failure_threshold", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.failure_threshold_expiration_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.max_address_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.device_fingerprint_challenge.0.max_address_count_expiration_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.action", "BLOCK"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.action_expiration_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_action", "SHOW_ERROR_PAGE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_error_page_code", "401"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_error_page_description", "blockErrorPageDescription2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_error_page_message", "blockErrorPageMessage2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_response_code", "401"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_footer", "captchaFooter2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_header", "captchaHeader2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_submit_label", "captchaSubmitLabel2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_title", "captchaTitle2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.failure_threshold", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.failure_threshold_expiration_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.interaction_threshold", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.is_nat_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.recording_period_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.set_http_header.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.set_http_header.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.human_interaction_challenge.0.set_http_header.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.action", "BLOCK"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.action_expiration_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.are_redirects_challenged", "true"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_action", "SHOW_ERROR_PAGE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_error_page_code", "401"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_error_page_description", "blockErrorPageDescription2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_error_page_message", "blockErrorPageMessage2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_response_code", "401"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_footer", "captchaFooter2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_header", "captchaHeader2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_submit_label", "captchaSubmitLabel2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_title", "captchaTitle2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.criteria.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.criteria.0.condition", "URL_STARTS_WITH"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.criteria.0.is_case_sensitive", "true"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.criteria.0.value", "/publ"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.failure_threshold", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.is_nat_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.set_http_header.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.set_http_header.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.js_challenge.0.set_http_header.0.value", "value2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.origin", "primary2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.origin_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.allowed_http_methods.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.block_action", "SHOW_ERROR_PAGE"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.block_error_page_code", "401"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.block_error_page_description", "blockErrorPageDescription2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.block_error_page_message", "blockErrorPageMessage2"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.block_response_code", "401"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.is_response_inspected", "false"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.max_argument_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.max_name_length_per_argument", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.max_response_size_in_ki_b", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.max_total_name_length_of_arguments", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.media_types.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.protection_settings.0.recommendations_period_in_days", "11"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.whitelists.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.whitelists.0.addresses.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "waf_config.0.whitelists.0.name", "name2"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_waas_waas_policies", "test_waas_policies", Optional, Update, waasPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + WaasPolicyResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_waas_waas_policy", "test_waas_policy", Optional, Update, waasPolicyRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_names.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ids.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "states.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_created_less_than"),

				resource.TestCheckResourceAttr(datasourceName, "waas_policies.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "waas_policies.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "waas_policies.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "waas_policies.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "waas_policies.0.domain", waasPolicyDomain),
				resource.TestCheckResourceAttr(datasourceName, "waas_policies.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "waas_policies.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "waas_policies.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "waas_policies.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_waas_waas_policy", "test_waas_policy", Required, Create, waasPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + WaasPolicyResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "waas_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "additional_domains.#", "2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cname"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "domain", waasPolicyDomain),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "origin_groups.#", "2"),
				CheckResourceSetContainsElementWithProperties(singularDatasourceName, "origin_groups", map[string]string{
					"origin_group.#": "2",
				},
					[]string{}),
				resource.TestCheckResourceAttr(singularDatasourceName, "origins.#", "2"),
				CheckResourceSetContainsElementWithProperties(singularDatasourceName, "origins", map[string]string{
					"custom_headers.#": "2",
					"http_port":        "80",
					"https_port":       "443",
					"uri":              "192.168.0.11",
				},
					[]string{}),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "policy_config.0.certificate_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.cipher_group", "DEFAULT"),

				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.health_checks.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.health_checks.0.expected_response_code_group.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.health_checks.0.expected_response_text", "expectedResponseText2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.health_checks.0.headers.%", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.health_checks.0.healthy_threshold", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.health_checks.0.interval_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.health_checks.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.health_checks.0.is_response_text_check_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.health_checks.0.method", "POST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.health_checks.0.path", "/"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.health_checks.0.timeout_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.health_checks.0.unhealthy_threshold", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.client_address_header", "X_FORWARDED_FOR"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.is_behind_cdn", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.is_cache_control_respected", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.is_https_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.is_https_forced", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.is_origin_compression_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.is_response_buffering_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.is_sni_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.load_balancing_method.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.load_balancing_method.0.domain", "example2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.load_balancing_method.0.expiration_time_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.load_balancing_method.0.method", "STICKY_COOKIE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.load_balancing_method.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.tls_protocols.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "policy_config.0.websocket_path_prefixes.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.action", "DETECT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.block_action", "SHOW_ERROR_PAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.block_error_page_code", "401"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.block_error_page_description", "blockErrorPageDescription2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.block_error_page_message", "blockErrorPageMessage2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.block_response_code", "401"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.bypass_challenges.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.captcha_footer", "captchaFooter2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.captcha_header", "captchaHeader2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.captcha_submit_label", "captchaSubmitLabel2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.captcha_title", "captchaTitle2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.criteria.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.criteria.0.condition", "URL_IS_NOT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.criteria.0.is_case_sensitive", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.criteria.0.value", "/secret"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.redirect_response_code", "MOVED_PERMANENTLY"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.redirect_url", "http://0.0.0.0:81"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.response_header_manipulation.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.response_header_manipulation.0.action", "ADD_HTTP_RESPONSE_HEADER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.response_header_manipulation.0.header", "header2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.access_rules.0.response_header_manipulation.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.address_rate_limiting.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.address_rate_limiting.0.allowed_rate_per_address", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.address_rate_limiting.0.block_response_code", "401"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.address_rate_limiting.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.address_rate_limiting.0.max_delayed_count_per_address", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.caching_rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.caching_rules.0.action", "CACHE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.caching_rules.0.caching_duration", "PT2S"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.caching_rules.0.client_caching_duration", "PT2S"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.caching_rules.0.criteria.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.caching_rules.0.criteria.0.condition", "URL_STARTS_WITH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.caching_rules.0.criteria.0.value", "/publ"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.caching_rules.0.is_client_caching_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.caching_rules.0.key", "key2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.caching_rules.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.captchas.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.captchas.0.failure_message", "failureMessage2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.captchas.0.footer_text", "footerText2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.captchas.0.header_text", "headerText2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.captchas.0.session_expiration_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.captchas.0.submit_label", "submitLabel2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.captchas.0.title", "title2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.captchas.0.url", "url2"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.custom_protection_rules.#", "1"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.custom_protection_rules.0.action", "BLOCK"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "waf_config.0.custom_protection_rules.0.id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.action", "BLOCK"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.action_expiration_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_action", "SHOW_ERROR_PAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_error_page_code", "401"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_error_page_description", "blockErrorPageDescription2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_error_page_message", "blockErrorPageMessage2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.block_response_code", "401"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_footer", "captchaFooter2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_header", "captchaHeader2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_submit_label", "captchaSubmitLabel2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.challenge_settings.0.captcha_title", "captchaTitle2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.failure_threshold", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.failure_threshold_expiration_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.max_address_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.device_fingerprint_challenge.0.max_address_count_expiration_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.action", "BLOCK"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.action_expiration_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_action", "SHOW_ERROR_PAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_error_page_code", "401"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_error_page_description", "blockErrorPageDescription2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_error_page_message", "blockErrorPageMessage2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.block_response_code", "401"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_footer", "captchaFooter2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_header", "captchaHeader2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_submit_label", "captchaSubmitLabel2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.challenge_settings.0.captcha_title", "captchaTitle2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.failure_threshold", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.failure_threshold_expiration_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.interaction_threshold", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.is_nat_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.recording_period_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.set_http_header.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.set_http_header.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.human_interaction_challenge.0.set_http_header.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.action", "BLOCK"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.action_expiration_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.are_redirects_challenged", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.challenge_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_action", "SHOW_ERROR_PAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_error_page_code", "401"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_error_page_description", "blockErrorPageDescription2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_error_page_message", "blockErrorPageMessage2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.challenge_settings.0.block_response_code", "401"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_footer", "captchaFooter2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_header", "captchaHeader2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_submit_label", "captchaSubmitLabel2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.challenge_settings.0.captcha_title", "captchaTitle2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.criteria.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.criteria.0.condition", "URL_STARTS_WITH"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.criteria.0.is_case_sensitive", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.criteria.0.value", "/publ"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.failure_threshold", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.is_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.is_nat_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.set_http_header.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.set_http_header.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.js_challenge.0.set_http_header.0.value", "value2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.origin", "primary2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.origin_groups.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.protection_settings.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.protection_settings.0.allowed_http_methods.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.protection_settings.0.block_action", "SHOW_ERROR_PAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.protection_settings.0.block_error_page_code", "401"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.protection_settings.0.block_error_page_description", "blockErrorPageDescription2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.protection_settings.0.block_error_page_message", "blockErrorPageMessage2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.protection_settings.0.block_response_code", "401"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.protection_settings.0.is_response_inspected", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.protection_settings.0.max_argument_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.protection_settings.0.max_name_length_per_argument", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.protection_settings.0.max_response_size_in_ki_b", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.protection_settings.0.max_total_name_length_of_arguments", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.protection_settings.0.media_types.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.protection_settings.0.recommendations_period_in_days", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.whitelists.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.whitelists.0.addresses.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "waf_config.0.whitelists.0.name", "name2"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + WaasPolicyResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckWaasWaasPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).waasClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_waas_waas_policy" {
			noResourceFound = false
			request := oci_waas.GetWaasPolicyRequest{}

			tmp := rs.Primary.ID
			request.WaasPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "waas")

			response, err := client.GetWaasPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_waas.WaasPolicyLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !InSweeperExcludeList("WaasWaasPolicy") {
		resource.AddTestSweepers("WaasWaasPolicy", &resource.Sweeper{
			Name:         "WaasWaasPolicy",
			Dependencies: DependencyGraph["waasPolicy"],
			F:            sweepWaasWaasPolicyResource,
		})
	}
}

func sweepWaasWaasPolicyResource(compartment string) error {
	waasClient := GetTestClients(&schema.ResourceData{}).waasClient()
	waasPolicyIds, err := getWaasPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, waasPolicyId := range waasPolicyIds {
		if ok := SweeperDefaultResourceId[waasPolicyId]; !ok {
			deleteWaasPolicyRequest := oci_waas.DeleteWaasPolicyRequest{}

			deleteWaasPolicyRequest.WaasPolicyId = &waasPolicyId

			deleteWaasPolicyRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "waas")
			_, error := waasClient.DeleteWaasPolicy(context.Background(), deleteWaasPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting WaasPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", waasPolicyId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &waasPolicyId, waasPolicySweepWaitCondition, time.Duration(3*time.Minute),
				waasPolicySweepResponseFetchOperation, "waas", true)
		}
	}
	return nil
}

func getWaasPolicyIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "WaasPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	waasClient := GetTestClients(&schema.ResourceData{}).waasClient()

	listWaasPoliciesRequest := oci_waas.ListWaasPoliciesRequest{}
	listWaasPoliciesRequest.CompartmentId = &compartmentId
	listWaasPoliciesResponse, err := waasClient.ListWaasPolicies(context.Background(), listWaasPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting WaasPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, waasPolicy := range listWaasPoliciesResponse.Items {
		id := *waasPolicy.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "WaasPolicyId", id)
	}
	return resourceIds, nil
}

func waasPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if waasPolicyResponse, ok := response.Response.(oci_waas.GetWaasPolicyResponse); ok {
		return waasPolicyResponse.LifecycleState != oci_waas.WaasPolicyLifecycleStateDeleted
	}
	return false
}

func waasPolicySweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.waasClient().GetWaasPolicy(context.Background(), oci_waas.GetWaasPolicyRequest{
		WaasPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
