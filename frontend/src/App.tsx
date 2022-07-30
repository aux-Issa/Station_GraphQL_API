import { BrowserRouter, Routes, Route } from "react-router-dom";
import logo from './logo.svg';
import './App.css';
import {GraphQLClient} from "graphql-request";
import {getSdk} from "./graphql/generated/graphql";
import Name from "./containers/Name";

function App() {

  // main()

  return (
    <BrowserRouter>
      <Routes>
        <Route path={`/name`} element={<Name />} />
        {/* <Route path={`/register/`} element={<AfterStation />} />
        <Route path={`/login/`} element={<BeforeStation />} /> */}
      </Routes>
    </BrowserRouter>
  );
}

export default App;
