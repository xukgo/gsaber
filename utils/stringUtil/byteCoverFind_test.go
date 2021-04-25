package stringUtil

import (
	"testing"
)

var sql_string1 = "CREATE TABLE `comp_0402` (\n  `id` int(11) NOT NULL AUTO_INCREMENT,\n  `session_id` varchar(64) NOT NULL,\n  `call_id` varchar(64) NOT NULL,\n  `host_name` varchar(64) NOT NULL,\n  `create_time` datetime DEFAULT NULL,\n  `progress_time` datetime DEFAULT NULL,\n  `answer_time` datetime DEFAULT NULL,\n  `hangup_time` datetime DEFAULT NULL,\n  `caller` varchar(32) NOT NULL,\n  `called` varchar(32) NOT NULL,\n  `display_name` varchar(32) NOT NULL,\n  `bill_number` varchar(32) NOT NULL,\n  `direction` int(1) NOT NULL DEFAULT '0',\n  `sip_cause` int(11) NOT NULL DEFAULT '0',\n  `soft_cause` int(11) NOT NULL DEFAULT '0',\n  `soft_description` varchar(30) DEFAULT NULL,\n  `disposition` varchar(32) NOT NULL,\n  `record_filename` varchar(255) NOT NULL,\n  `record_expiretime` datetime DEFAULT NULL,\n  `call_from` varchar(32) NOT NULL,\n  `local_sip_addr` varchar(32) NOT NULL,\n  `local_sip_port` int(11) NOT NULL,\n  `remote_sip_addr` varchar(32) NOT NULL,\n  `remote_sip_port` int(11) NOT NULL,\n  `trunk_name` varchar(32) NOT NULL,\n  `gateway_name` varchar(64) NOT NULL,\n  `local_media_ip` varchar(32) NOT NULL,\n  `local_media_port` int(11) NOT NULL,\n  `remote_media_ip` varchar(32) NOT NULL,\n  `remote_media_port` int(11) NOT NULL,\n  `codec` varchar(32) NOT NULL,\n  `recv_packet` int(11) NOT NULL DEFAULT '0',\n  `send_packet` int(11) NOT NULL DEFAULT '0',\n  `recv_bytes` int(11) NOT NULL DEFAULT '0',\n  `send_bytes` int(11) NOT NULL DEFAULT '0',\n  `lost_packet` int(11) NOT NULL DEFAULT '0',\n  `min_jitter` int(11) NOT NULL,\n  `max_jitter` int(11) NOT NULL,\n  `call_duration` int(11) NOT NULL,\n  `bill_duration` int(11) NOT NULL,\n  `pdd_msec` int(11) NOT NULL,\n  `ring_sec` int(11) NOT NULL,\n  `notify_callback` varchar(255) NOT NULL,\n  `customer_id` int(11) NOT NULL,\n  `customer_name` varchar(255) NOT NULL,\n  `user_id` int(11) NOT NULL,\n  `user_identify` varchar(32) DEFAULT NULL,\n  `sub_user_id` int(11) DEFAULT NULL,\n  `account_id` int(11) NOT NULL,\n  `app_id` int(11) NOT NULL,\n  `app_identify` varchar(32) DEFAULT NULL,\n  `app_name` varchar(255) NOT NULL,\n  `is_agent` int(11) NOT NULL DEFAULT '0',\n  `charge_mode` int(11) NOT NULL DEFAULT '0',\n  `customer_source` int(11) NOT NULL DEFAULT '0',\n  `industry_info` int(11) DEFAULT NULL,\n  `customer_area` int(11) DEFAULT NULL,\n  `rate_package_id` int(11) DEFAULT NULL,\n  `rate_package_name` varchar(255) DEFAULT NULL,\n  `called_type` int(11) DEFAULT NULL,\n  `business_capability` int(11) DEFAULT NULL,\n  `business_capability_name` varchar(32) DEFAULT NULL,\n  `json_data` varchar(255) DEFAULT NULL,\n  `dtmf` varchar(255) DEFAULT NULL COMMENT '按键码',\n  `title` varchar(64) DEFAULT NULL COMMENT '模板名称',\n  `template_id` varchar(32) DEFAULT NULL COMMENT '模板id',\n  `called_area` int(11) DEFAULT NULL,\n  `carrier` int(11) DEFAULT NULL,\n  `time_profile_type` int(11) DEFAULT NULL,\n  `fee_type` int(11) DEFAULT NULL,\n  `rate` decimal(11,4) DEFAULT NULL,\n  `rate_unit` int(11) DEFAULT NULL,\n  `rate_type` int(11) DEFAULT NULL,\n  `rate_interval` int(11) DEFAULT NULL,\n  `cost_total` decimal(11,4) DEFAULT NULL,\n  `cost_detail` text,\n  `time` datetime NOT NULL,\n  PRIMARY KEY (`id`),\n  UNIQUE KEY `index_comp_0402_call_id` (`call_id`),\n  KEY `index_comp_0402_ui_time` (`user_identify`,`time`),\n  KEY `index_comp_0402_json_data` (`json_data`),\n  KEY `index_comp_0402_template_id` (`template_id`),\n  KEY `index_comp_0402_time` (`time`),\n  KEY `index_comp_0402_app_id` (`app_id`),\n  KEY `index_comp_0402_at_ct` (`answer_time`,`create_time`),\n  KEY `index_comp_0402_at_create_time` (`create_time`),\n  KEY `index_comp_0402_at_bc_direction_caller` (`business_capability`,`direction`,`caller`),\n  KEY `index_comp_0402_at_bc_direction_bn` (`business_capability`,`direction`,`bill_number`),\n  KEY `index_comp_0402_at_bc_direction_called` (`business_capability`,`direction`,`called`)\n)"

func Test_coverFind0(t *testing.T) {
	str := sql_string1
	finder := new(ByteCoverFinder)
	finder.AddCover(InitByteCover('(', ')'))
	finder.AddCover(InitByteCover('`', '`'))
	buff := []byte(str)
	index := finder.Index(buff, len(buff), ')')
	if index != -1 {
		t.FailNow()
	}
}
func Test_coverFind1(t *testing.T) {
	str := "`rate` decimal(11,4) DEFAULT NULL,"
	finder := new(ByteCoverFinder)
	finder.AddCover(InitByteCover('(', ')'))
	finder.AddCover(InitByteCover('`', '`'))
	buff := []byte(str)
	index := finder.Index(buff, len(buff), ',')
	if index != len(buff)-1 {
		t.FailNow()
	}
	index = finder.Index(buff, len(buff), ')')
	if index != -1 {
		t.FailNow()
	}
}
func Test_coverFind2(t *testing.T) {
	str := "1+(d+'var1*(b*c)')*(a+(b*c/(d+f)))@"
	finder := new(ByteCoverFinder)
	finder.AddCover(InitByteCover('(', ')'))
	finder.AddCover(InitByteCover(0x27, 0x27))
	buff := []byte(str)
	index := finder.Index(buff, len(buff), '@')
	if index != len(buff)-1 {
		t.FailNow()
	}
	index = finder.Index(buff, len(buff), ')')
	if index != -1 {
		t.FailNow()
	}
}

func Test_GetFirstCompareSegment(t *testing.T) {
	str := "1+(d+'var1*(b*c)')*(a+(b*c/(d+f)))"
	finder := new(ByteCoverFinder)
	finder.AddCover(InitByteCover('(', ')'))
	finder.AddCover(InitByteCover(0x27, 0x27))
	buff := []byte(str)
	seg := finder.GetFirstSegment(buff, len(buff), '(', ')')
	if len(seg) != 16 {
		t.FailNow()
	}
	segs := finder.GetSegments(buff, len(buff), '(', ')')
	if len(segs) != 2 {
		t.FailNow()
	}
}
