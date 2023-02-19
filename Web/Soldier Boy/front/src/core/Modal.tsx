'use client'

import {
  createContext,
  useContext,
  useState,
  useMemo,
  useCallback,
  Fragment,
} from 'react'
import { useSubmitFlag, useTask } from './hooks'

interface IContext {
  setter: (obj: any) => void
  close: () => void
}

const Context = createContext<IContext>({
  setter: () => {},
  close: () => {},
})

export function useModal() {
  return useContext(Context)
}

export default function Modal({ children }: { children: React.ReactNode }) {
  const task = useTask()
  const [open, setOpen] = useState(false)

  const flag = useSubmitFlag()

  const setter = useCallback((id: number) => {
    task.loadTask(id)
    setOpen(true)
  }, [])

  const close = useCallback(() => {
    setOpen(false)
    task.clear()
    flag.clear()
  }, [])

  const value = useMemo(
    () => ({
      setter,
      close,
    }),
    [setter, close]
  )

  return (
    <Context.Provider value={value}>
      {open ? (
        <div className="fixed inset-0 z-50 flex items-center justify-center">
          <div
            className="absolute inset-0 bg-black opacity-50 z-10"
            onClick={close}
          ></div>
          <div className="bg-[#333] rounded-lg shadow-lg p-10 z-20 w-[500px]">
            {task.isLoading || !task.data ? (
              <div className="grid place-content-center ">
                <h1 className="text-center text-xl type02">loading..</h1>
              </div>
            ) : (
              <div>
                <div className="pb-4">
                  <h1 className="text-center font-black text-2xl type02">
                    {task.data.Title}
                  </h1>
                </div>
                <p className="test-xs">
                  {task.data.Description.split('\n').map(l => (
                    <Fragment key={l}>
                      {l}
                      <br />
                    </Fragment>
                  ))}
                </p>

                <div className="py-4 flex flex-col">
                    <div>
                      <a href={task.data!.Url} target="_blank" className="text-blue-500 ">
                        {task.data!.Url}
                      </a>
                    </div>
                </div>

                  <form onSubmit={flag.submit(task.data.ID)}>
                    <input
                      disabled={flag.isLoading}
                      type="text"
                      placeholder="Enter Flag.."
                      name="flag"
                      required
                      className="bg-[#222324] text-white focus:outline-none rounded-lg px-3 py-2 w-full mt-1 disabled:opacity-50"
                    />

                    {flag.submitted ? (
                      <div
                        className={`mt-4 text-white px-5 py-3 rounded-lg font-medium ${
                          flag.correct ? 'bg-green-500' : 'bg-red-500'
                        }`}
                      >
                        {flag.correct ? 'Correct!' : 'Wrong!'}
                      </div>
                    ) : null}
                  </form>
              </div>
            )}
          </div>
        </div>
      ) : null}
      {children}
    </Context.Provider>
  )
}
