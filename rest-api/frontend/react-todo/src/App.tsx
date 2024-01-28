import { useEffect } from 'react'
import axios from 'axios'   // APIへのリクエスト用

import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { Auth } from './components/Auth'
import { Todo } from './components/Todo'
import { CsrfToken } from './types';


function App() {
  useEffect( () => {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      // GET通信で /csrf へリクエスト
      const { data } = await axios.get<CsrfToken>(
        `${process.env.REACT_APP_API_URL}/csrf`
      )
      // 得たcsrf_token を header へ格納
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }

    getCsrfToken()
  }, [])

  return (
    <BrowserRouter>
      <Routes>
          <Route path = "/" element = {<Auth />} />
          <Route path = "/todo" element = {<Todo />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
