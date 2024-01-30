
import axios from 'axios'   // APIへのリクエスト用
import useStore from '../store'
import { Credential } from '../types'
import { useError } from './useError'

import { useNavigate } from 'react-router-dom'    // 画面遷移用
import { useMutation } from '@tanstack/react-query'   // 値変更を検知する


export const useMutateAuth = () => {
    const navigate = useNavigate()
    const resetEditedTask = useStore((state) => state.resetEditedTask)
    const { switchErrorHandling } = useError()

    const loginMutation = useMutation(
        async (user: Credential) => await axios.post(`${process.env.REACT_APP_API_URL}/login`, user), {
            onSuccess: () => {
                navigate('/todo')
            },
            onError: (err: any) => {
                if (err.response.data.message) {
                    switchErrorHandling(err.response.data.message)
                } else {
                    switchErrorHandling(err.response.data)
                }
            }
        }
    )

    const registerMutation = useMutation(
        async (user: Credential) => await axios.post(`${process.env.REACT_APP_API_URL}/signup`, user), {
            // TODO: onSuccess の処理を追加

            onError: (err: any) => {
                if (err.response.data.message) {
                    switchErrorHandling(err.response.data.message)
                } else {
                    switchErrorHandling(err.response.data)
                }
            }
        }
    )

    const logoutMutation = useMutation(
        async () => await axios.post(`${process.env.REACT_APP_API_URL}/logout`), {
            onSuccess: () => {
                resetEditedTask()
                navigate('/')
            },
            onError: (err: any) => {
                if (err.response.data.message) {
                    switchErrorHandling(err.response.data.message)
                } else {
                    switchErrorHandling(err.response.data)
                }
            },
        }
    )

    return { loginMutation, registerMutation, logoutMutation }
}




