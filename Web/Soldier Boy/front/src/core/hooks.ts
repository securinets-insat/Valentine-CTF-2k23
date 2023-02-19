import { Task } from '@/@types'
import axios from './api'
import { useState, useEffect, useCallback } from 'react'

export function useTasks() {
  const [tasks, setTasks] = useState<Task[]>([])
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    axios.get('/tasks').then(response => {
      setTasks(response.data)
      setIsLoading(false)
    })
  }, [])

  return {
    tasks,
    isLoading,
  }
}

export function useTask() {
  const [data, setTask] = useState<Task | null>(null)
  const [isLoading, setIsLoading] = useState(true)

  const loadTask = useCallback(
    (id: number) => {
      setIsLoading(true)
      axios.get(`/tasks/${id}`).then(response => {
        setTask(response.data)
        setIsLoading(false)
      })
    },
    [setTask, setIsLoading]
  )

  const clear = useCallback(() => {
    setTask(null)
  }, [setTask])

  return {
    clear,
    data,
    isLoading,
    loadTask,
  }
}

export function useSubmitFlag() {
  const [isLoading, setIsLoading] = useState(false)
  const [correct, setCorrect] = useState(false)
  const [submitted, setSubmitted] = useState(false)

  const submit = useCallback(
    (id: number) => {
      return (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        setSubmitted(false)
        const formData = new FormData(e.currentTarget)
        const flag = formData.get('flag') as string
        setIsLoading(true)
        return axios
          .post(`/tasks/${id}/submit`, { flag })
          .then(response => {
            setCorrect(response.data === 'Correct flag')
            setIsLoading(false)
            setSubmitted(true)
          })
          .catch(() => {
            setCorrect(false)
            setIsLoading(false)
            setSubmitted(true)
          })
      }
    },

    [setIsLoading, setCorrect, setSubmitted]
  )

  const clear = useCallback(() => {
    setCorrect(false)
    setSubmitted(false)
  }, [setCorrect, setSubmitted])

  return {
    submit,
    isLoading,
    correct,
    submitted,
    clear,
  }
}
