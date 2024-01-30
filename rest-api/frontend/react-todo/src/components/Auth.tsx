import { useState, FormEvent } from 'react'
import { CheckBadgeIcon, ArrowPathIcon } from '@heroicons/react/24/solid'
import { useMutateAuth } from '../hooks/useMutateAuth'


export const Auth = () => {
    //     ↓変数 ↓変更用メソッド  ↓ useStateで値変更検知を可能にしている
    const [email, setEmail] = useState('')
    const [pw, setPw] = useState('')
    const [isLogin, setIsLogin] = useState(true)
    const { loginMutation, registerMutation } = useMutateAuth()

    const submitAuthHandler = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        if (isLogin) {
            // Login処理を呼び出す
            loginMutation.mutate({
                email: email,
                password: pw,
            })
        } else {
            // Signup + Login
            await registerMutation
                .mutateAsync({
                    email: email,
                    password: pw,
                })
                .then(() => loginMutation.mutate({
                    email: email,
                    password: pw,
                }))
        }
    }

    return (
        <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
          {/* タイトル */}
          <div className="flex items-center">
            <CheckBadgeIcon className="h-8 w-8 mr-2 text-blue-500" />
            <span className="text-center text-3xl font-extrabold">
              Todo app by React/Go(Echo)
            </span>
          </div>

          <h2 className="my-6">{isLogin ? 'Login' : 'Create a new account'}</h2>

          <form onSubmit={submitAuthHandler}>
            {/* Email 入力 */}
            <div>
              <input
                className="mb-3 px-3 text-sm py-2 border border-gray-300"
                name="email"
                type="email"
                autoFocus
                placeholder="Email address"
                onChange={(e) => setEmail(e.target.value)}
                value={email}
              />
            </div>

            {/* Password 入力 */}
            <div>
              <input
                className="mb-3 px-3 text-sm py-2 border border-gray-300"
                name="password"
                type="password"
                placeholder="Password"
                onChange={(e) => setPw(e.target.value)}
                value={pw}
              />
            </div>

            {/* Login/Signup ボタン */}
            <div className="flex justify-center my-2">
              <button
                className="disabled:opacity-40 py-2 px-4 rounded text-white bg-indigo-600"
                disabled={!email || !pw}
                type="submit"
              >
                {isLogin ? 'Login' : 'Sign Up'}
              </button>
            </div>
          </form>

          <ArrowPathIcon
            onClick={() => setIsLogin(!isLogin)}
            className="h-6 w-6 my-2 text-blue-500 cursor-pointer"
          />
        </div>
      )
}
