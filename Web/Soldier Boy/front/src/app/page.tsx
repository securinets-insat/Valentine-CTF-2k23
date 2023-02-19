'use client'
import Task from '@/core/Task'
import Modal from '@/core/Modal'
import { useTasks } from '@/core/hooks'

export default function Home() {
  const { isLoading, tasks } = useTasks()

  if (isLoading) {
    return (
      <div className="grid place-content-center h-screen">
        <h1 className="text-center text-5xl type02">loading..</h1>
      </div>
    )
  }

  return (
    <Modal>
      <main>
        <div className="py-20">
          <div className="type02 text-center flex flex-col">
            <h1 className="text-7xl">hmlndr</h1>
            <span className="text-xs">from</span>
            <h2 className="text-3xl">nixmnd</h2>
            <h3 className="tracking-[0.5rem]">ctf collection</h3>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 gap-10 px-4 sm:px-10 md:px-20 xl:px-40 py-20">
            {tasks.map(task => (
              <Task
                key={task.ID}
                id={task.ID}
                title={task.Title}
              />
            ))}
          </div>
        </div>
      </main>
    </Modal>
  )
}
