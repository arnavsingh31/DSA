package hashmaps

// Refer:- https://leetcode.com/problems/isomorphic-strings/editorial/?envType=study-plan-v2&envId=top-interview-150 for better understanding.

func IsIsomorphic(s string, t string) bool {
	hashMap_s_t := make(map[string]string)
	hashMap_t_s := make(map[string]string)

	for i := 0; i < len(s); i++ {
		val_s_t, exist_s_t := hashMap_s_t[string(s[i])]
		val_t_s, exist_t_s := hashMap_t_s[string(t[i])]

		if !exist_s_t && !exist_t_s {
			// If key/char doesn't exist in either map then we add corresponding key to Map
			hashMap_s_t[string(s[i])] = string(t[i])
			hashMap_t_s[string(t[i])] = string(s[i])

		} else if (exist_s_t) && (val_s_t != string(t[i])) {
			// if key exist in map of S->T but the value/char isn't same as existing one then we return false.
			return false

		} else if (exist_t_s) && (val_t_s != string(s[i])) {
			// if key exist in map of T->S but the value/char isn't same as existing one then we return false.
			return false
		}
	}
	return true
}
