'use client'
import { useModal } from './Modal'

export default function Task(props: TaskProps) {
  const m = useModal()

  function onClick() {
    m.setter(props.id)
  }

  return (
    <button
      className={`p-10 rounded-xl bg-[#222324]`}
      onClick={onClick}
    >
      <h1 className="text-center font-black text-xl type02">{props.title}</h1>
    </button>
  )
}

interface TaskProps {
  id: number
  title: string
}
