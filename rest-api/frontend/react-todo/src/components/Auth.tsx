import { useState, FormEvent } from 'react'
import { CheckBadgeIcon, ArrowPathIcon } from '@heroicons/react/24/solid'
import { useMutateAuth } from '../hooks/useMutateAuth'


export const Auth = () => {
    //     ↓変数 ↓変更用メソッド  ↓ useStateで値変更検知を可能にしている
    const [email, setEmail] = useState('')
    const [pw, setPw] = useState('')
    const [isLogin, setIsLogin] = useState(true)
    const { loginMutation, registerMutation } = useMutateAuth()

    const submitAuthHadler = async (e: FormEvent<HTMLFormElement>) => {
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

    return <div>Auth</div>
}
