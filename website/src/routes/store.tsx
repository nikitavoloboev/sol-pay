import { Match, Show, Switch, createSignal } from "solid-js"
// import Bought from "~/components/Bought"
import Card from "~/components/Card"
import Info from "~/components/Info"
import Pricing from "~/components/Pricing"

export default function Store() {
  // doesnt save if u have bought that course
  const [showInfoCard, setShowInfoCard] = createSignal("")
  return (
    <>
      <style>
        {`
      #InfoCard {
        animation: 0.2s InfoCardSlide forwards ease-out
      }
      @keyframes InfoCardSlide {
        0% {
          transform: translateX(850px)
        }
      }
      `}
      </style>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2 p-2">
        <div class="fixed top-3 left-0 w-full flex items-center justify-between px-4">
          <div class="flex gap-2">
            <div class="bg-white p-0.5 px-1 gap-2 items-center flex justify-center rounded-[6px] border-slate-500 border border-opacity-50 font-semibold text-black text-opacity-70">
              <div class="px-[16px] hover:bg-neutral-300 rounded-[6px] p-[8px]">
                Products
              </div>
              <div class="px-[16px] hover:bg-neutral-300 rounded-[6px] p-[8px]">
                Ideas
              </div>
            </div>
          </div>
          <div class="flex gap-2 items-center h-full">
            <div class="bg-white h-[46px] items-center flex justify-center rounded-[6px] w-[150px] p-3 border-slate-500 border border-opacity-50 hover:bg-black hover:text-white transition-all font-semibold text-black text-opacity-70">
              Sign In
            </div>
          </div>
        </div>

        <Card setShowInfoCard={setShowInfoCard} />
        <Card setShowInfoCard={setShowInfoCard} />
        <Card setShowInfoCard={setShowInfoCard} />
        <Card setShowInfoCard={setShowInfoCard} />
        <Card setShowInfoCard={setShowInfoCard} />
        <Card setShowInfoCard={setShowInfoCard} />
        <Show when={showInfoCard()}>
          <div class="fixed top-0 left-0 w-screen h-screen flex items-center justify-end">
            <div
              class="w-screen h-screen fixed top-0 left-0 bg-[#f4f1f0] bg-opacity-20 backdrop-blur-sm"
              onClick={() => {
                setShowInfoCard("")
              }}
            ></div>
            <div
              id="InfoCard"
              class=" relative w-[850px] z-20 h-full p-2 gap-2 flex bg-neutral-200"
              style={{ "border-radius": "6px 0 0 6px" }}
            >
              <Switch>
                <Match when={showInfoCard() === "Pricing"}>
                  <Pricing setShowInfoCard={setShowInfoCard}></Pricing>
                </Match>
                <Match when={showInfoCard() === "Info"}>
                  <Info setShowInfoCard={setShowInfoCard}></Info>
                </Match>
                <Match when={showInfoCard() === "Bought"}>
                  {/* <Bought></Bought> */}
                </Match>
              </Switch>
            </div>
          </div>
        </Show>
      </div>
    </>
  )
}
