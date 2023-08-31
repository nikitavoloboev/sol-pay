import { createSignal } from "solid-js"

export default function Nav() {
  const [name, setName] = createSignal("Anatoly")
  return (
    <div class="w-full p-8 px-[10%] flex justify-between items-center border-b-2 border-black">
      <div class="font-light text-[20px]">{name()}</div>
      <div class="bg-black rounded-[4px] border-black border hover:scale-[1.1] text-white px-[12px] p-[8px]">
        Sign In
      </div>
    </div>
  )
}
