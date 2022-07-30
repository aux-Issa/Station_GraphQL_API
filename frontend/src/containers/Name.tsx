import { debug } from "console";
import {GraphQLClient} from "graphql-request";
import {getSdk} from "../graphql/generated/graphql";

function Name() {
  async function main() {
    const endpoint = "http://localhost:8080/graphql"
    const client = new GraphQLClient(endpoint)
    const sdk = getSdk(client)
  
    const res = await sdk.stationByCD(
      {
        id: 1130201,
      }
    )
    console.log(res)
    debugger
    return res
  }
  // const res = main()
  return(    
    <>
      <h1>新規登録ページ</h1>
      <div>
        {/* {res.stationByCD.lineName} */}
      </div>
      <div>
        {/* <Link to={`/`}>ホームに戻る</Link> */}
      </div>
    </>
  )
}

export default Name