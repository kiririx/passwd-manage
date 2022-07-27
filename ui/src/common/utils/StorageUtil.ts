import StringUtil from "./StringUtil";

export default class StorageUtil {
    public static get = (key: string) =>  {
        let v = sessionStorage.getItem(key)
        if (StringUtil.isBlank(v)) {
            v = localStorage.getItem(key)
        }
        return v
    }
}