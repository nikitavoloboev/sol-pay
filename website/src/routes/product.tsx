import { Match, Switch, createSignal } from "solid-js"
import FrontPage from "~/components/FrontPage"
import Product from "~/components/Product"

export default function product() {
  const [currentTab, setCurrentTab] = createSignal("FrontPage")
  return (
    <>
      <style>
        {`
      #Active {

        color: rgba(34, 197, 94);
      }
      #Inactive {

      }
      `}
      </style>
      <div class="flex items-center w-screen h-screen text-black dark:text-white">
        <div class="w-[15%] flex flex-col dark:bg-black bg-gray-200 h-full">
          <div class="border-b border-slate-400 border-opacity-30 h-[100px] flex items-center font-light text-[24px] justify-center">
            Main
          </div>
          <div>
            <div
              id={currentTab() === "FrontPage" ? "Active" : "Inactive"}
              class="p-4 px-8 border-b border-slate-400 border-opacity-30 font-light text-[16px] opacity-60 "
              onClick={() => {
                setCurrentTab("FrontPage")
              }}
            >
              Home
            </div>
            <div
              id={currentTab() === "Product" ? "Active" : "Inactive"}
              class="p-4 px-8 border-b border-slate-400 border-opacity-30  font-light text-[16px] opacity-60"
              onClick={() => {
                setCurrentTab("Product")
              }}
            >
              Product
            </div>
          </div>
        </div>
        <Switch>
          <Match when={currentTab() === "FrontPage"}>
            <FrontPage />
          </Match>
          <Match when={currentTab() === "Product"}>
            <Product />
          </Match>
        </Switch>
      </div>
    </>
  )
}
