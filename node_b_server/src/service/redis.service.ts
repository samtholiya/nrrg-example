import { IServicePanel } from "../namespace/initializer";
import * as redis from "redis"

@IServicePanel.register
export class RedisService {
    private static client: redis.RedisClient
    public init() {
        RedisService.client = redis.createClient(process.env["REDIS_URL"])
    }

    public static GetValueAndDelete(key: string, callback: redis.Callback<string>) {
        RedisService.client.get(key, (err, reply)=>{
            callback(err, reply);
            RedisService.client.del(key);
        })
    }
}