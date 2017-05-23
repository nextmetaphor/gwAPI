## Useful Tyk Redis Information

####Analytics Information
This is extracted (and removed from Redis) by the Tyk Pump component
```
> llen analytics-tyk-system-analytics
(integer) 10

> lrange analytics-tyk-system-analytics 0 10
```

####Get API Keys
```
> keys apikey*
1) "apikey-9088e38ae7d1447a70511482d3ef76e3"

> get apikey-9088e38ae7d1447a70511482d3ef76e3
"{\"last_check\":0,\"allowance\":1000,\"rate\":1000,\"per\":60,\"expires\":1,\"quota_max\":1,\"quota_renews\":1492076052,\"quota_remaining\":0,\"quota_renewal_rate\":60,\"access_rights\":{\"1\":{\"api_name\":\"my-api\",\"api_id\":\"1\",\"versions\":[\"v1\"],\"allowed_urls\":null}},\"org_id\":\"\",\"oauth_client_id\":\"\",\"oauth_keys\":{},\"basic_auth_data\":{\"password\":\"\",\"hash_type\":\"\"},\"jwt_data\":{\"secret\":\"\"},\"hmac_enabled\":true,\"hmac_string\":\"MTZiOTA2Nzg2Y2NmNDE0OTVhNzIxNzFlYWM4NDUyYjg=\",\"is_inactive\":true,\"apply_policy_id\":\"\",\"data_expires\":0,\"monitor\":{\"trigger_limits\":[]},\"enable_detail_recording\":false,\"meta_data\":{},\"tags\":[],\"alias\":\"\",\"last_updated\":\"1492075992\",\"id_extractor_deadline\":0,\"session_lifetime\":0}"
```