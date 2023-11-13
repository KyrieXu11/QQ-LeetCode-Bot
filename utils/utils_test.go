package utils

import "testing"

func TestDict_Get(t *testing.T) {
	cookies := "gr_user_id=92a8b27b-9f0d-4648-a0a0-4233e9edb41c; _bl_uid=OXlk3o8eiIaqaX6LgmemaIywt409; a2873925c34ecbd2_gr_last_sent_cs1=kyriexu; _gid=GA1.2.1272090214.1699616983; csrftoken=OzO3z9vTv6mCD5CMd6MK25RMJN8agfdGQqFbnOwIuzkjkUosm7is7V5K6LSlqso9; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMjg3NTM1IiwiX2F1dGhfdXNlcl9iYWNrZW5kIjoiYWxsYXV0aC5hY2NvdW50LmF1dGhfYmFja2VuZHMuQXV0aGVudGljYXRpb25CYWNrZW5kIiwiX2F1dGhfdXNlcl9oYXNoIjoiZmJjZjA0MTY5NjY0MTcyNThjYmU3MzlkYjA2NDQ3ZTdlMWI3NTcxOWI3ZDRlOWVmYjhlZjUzZmQzZTIxMGYyOSIsImlkIjoyODc1MzUsImVtYWlsIjoia3lyaWV4dTExQDEyNi5jb20iLCJ1c2VybmFtZSI6Imt5cmlleHUiLCJ1c2VyX3NsdWciOiJreXJpZXh1IiwiYXZhdGFyIjoiaHR0cHM6Ly9hc3NldHMubGVldGNvZGUuY24vYWxpeXVuLWxjLXVwbG9hZC9kZWZhdWx0X2F2YXRhci5wbmciLCJwaG9uZV92ZXJpZmllZCI6dHJ1ZSwiX3RpbWVzdGFtcCI6MTY5OTYxNzAxMS4wNDEyNzA3LCJleHBpcmVkX3RpbWVfIjoxNzAyMTQ4NDAwLCJ2ZXJzaW9uX2tleV8iOjB9.vHz0SSbIQ05Dvlm8cxhLh6xrAWr3dICysOhHEK5rHsQ; Hm_lvt_f0faad39bcf8471e3ab3ef70125152c3=1699153115,1699616983,1699625008,1699627355; _gat=1; a2873925c34ecbd2_gr_session_id=aa142ba7-1af4-42c3-b039-b6af0b4c3c91; a2873925c34ecbd2_gr_last_sent_sid_with_cs1=aa142ba7-1af4-42c3-b039-b6af0b4c3c91; a2873925c34ecbd2_gr_session_id_sent_vst=aa142ba7-1af4-42c3-b039-b6af0b4c3c91; _ga=GA1.1.1874174717.1659854208; Hm_lpvt_f0faad39bcf8471e3ab3ef70125152c3=1699627357; a2873925c34ecbd2_gr_cs1=kyriexu; _ga_PDVPZYN3CW=GS1.1.1699627354.25.1.1699627359.55.0.0"
	res := ParseCookies(cookies)
	t.Log(res)
}