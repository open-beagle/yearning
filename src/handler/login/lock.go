package login

import "time"

type loginStatus struct {
	Username string    `json:"username"`
	Time     time.Time `json:"time"`
	Count    int       `json:"count"`
}

var loginStatusMap = make(map[string]*loginStatus)

func LoginStartCheck(u loginForm) bool {
	lgSus, ok := loginStatusMap[u.Username]
	// 多次失败
	dur30s := time.Second * 30
	now := time.Now()
	if ok && lgSus.Count >= 3 {
		if (now.Sub(lgSus.Time)) > dur30s*10 {
			// 5分钟惩罚时间不允许登录，已超过，允许登录
			delete(loginStatusMap, lgSus.Username)
		} else {
			// 5分钟惩罚时间不允许登录
			return true
		}
	}
	return false
}

func LoginEndCheck(u loginForm, b bool) bool {
	lgSus, ok := loginStatusMap[u.Username]
	if b {
		if ok {
			// 登录成功，且存在锁，此时移除锁
			delete(loginStatusMap, lgSus.Username)
		}
	} else {
		// 以下是登录失败的处理逻辑
		if !ok {
			// 初次失败
			item := &loginStatus{
				Username: u.Username,
				Time:     time.Now(),
				Count:    1,
			}
			loginStatusMap[u.Username] = item
		} else {
			// 多次失败
			dur30s := time.Second * 30
			now := time.Now()
			if (now.Sub(lgSus.Time)) > dur30s {
				item := &loginStatus{
					Username: u.Username,
					Time:     time.Now(),
					Count:    1,
				}
				loginStatusMap[u.Username] = item
			} else {
				lgSus.Count = lgSus.Count + 1
				lgSus.Time = now
				if lgSus.Count >= 3 {
					// 短时间重复登录3次出现密码错误，准备封禁用户
					return true
				}
			}
		}
	}
	return false
}
