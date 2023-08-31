import { createSignal } from "solid-js"

export default function Nav() {
  const [name, setName] = createSignal("Anatoly")
  return (
    <div>
      <div>{name()}</div>
    </div>
  )
}
