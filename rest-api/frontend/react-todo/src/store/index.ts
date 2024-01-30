import { create } from 'zustand'   // 状態管理用

type EditedTask = {
    id: number
    title: string
}

type State = {
    // 状態管理したいデータ
    editedTask: EditedTask

    // 状態を変更する際のメソッド
    updateEditedTask: (payload: EditedTask) => void
    resetEditedTask: () => void
}

const useStore = create<State>((set) => ({
    editedTask: { id: 0, title: '' },

    updateEditedTask: (payload) => set({
        editedTask: payload
    }),
    resetEditedTask: () => set({ editedTask: { id: 0, title: '' }}),
}))

export default useStore