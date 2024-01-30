import axios from 'axios'
import { useQueryClient, useMutation } from '@tanstack/react-query'
import { Task } from '../types'
import useStore from '../store'
import { useError } from '../hooks/useError'

export const useMutateTask = () => {
    const queryClient = useQueryClient()
    const { switchErrorHandling } = useError()
    const resetEditedTask = useStore((state) => state.resetEditedTask)

    const createTaskMutation = useMutation(
        (task: Omit<Task, 'id' | 'created_at' | 'updated_at'>) => axios.post<Task>(`${process.env.REACT_APP_API_URL}/tasks`, task), {
            onSuccess: (res) => {
                // キャッシュの tasksデータ を確認
                const previsiousTasks = queryClient.getQueryData<Task[]>(['tasks'])
                // もしキャッシュがある時は
                if (previsiousTasks) {
                    // キャッシュを更新
                    queryClient.setQueryData(['tasks'], [...previsiousTasks, res.data])
                }
                resetEditedTask()
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

    const updateTaskMutation = useMutation(
        (task: Omit<Task, 'created_at' | 'updated_at'>) => axios.put<Task>(`${process.env.REACT_APP_API_URL}/tasks/${task.id}`, {
            title: task.title,
        }), {
            onSuccess: (res, variables) => {
                // キャッシュの tasksデータ を確認
                const previsiousTasks = queryClient.getQueryData<Task[]>(['tasks'])
                // もしキャッシュがある時は
                if (previsiousTasks) {
                    // キャッシュを更新
                    queryClient.setQueryData<Task[]>(
                        ['tasks'],
                        previsiousTasks.map((task) => task.id === variables.id ? res.data : task)
                    )
                }
                resetEditedTask()
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

    const deleteTaskMutation = useMutation(
        (id: number) => axios.delete<Task>(`${process.env.REACT_APP_API_URL}/tasks/${id}`), {
            onSuccess: (_, variables) => {
                // キャッシュの tasksデータ を確認
                const previsiousTasks = queryClient.getQueryData<Task[]>(['tasks'])
                // もしキャッシュがある時は
                if (previsiousTasks) {
                    // キャッシュを更新
                    queryClient.setQueryData<Task[]>(
                        ['tasks'],
                        previsiousTasks.filter((task) => task.id !== variables)
                    )
                }
                resetEditedTask()
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

    return {
        createTaskMutation,
        updateTaskMutation,
        deleteTaskMutation,
    }
}
