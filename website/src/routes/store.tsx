import { Show, createSignal } from "solid-js"
import Card from "~/components/Card"

export default function store() {
  const [showInfoCard, setShowInfoCard] = createSignal(true)
  return (
    <>
      <style>
        {`
      #InfoCard {
        animation: 0.2s InfoCardSlide forwards linear
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
            <div class="p-4 h-[54px] w-[54px] bg-blue-700 border-slate-500 border border-opacity-50 rounded-[6px]"></div>
            <div class="bg-white p-0.5 px-1 gap-2 items-center flex justify-center rounded-[6px] border-slate-500 border border-opacity-50 font-semibold text-black text-opacity-70">
              <div class="px-[16px] hover:bg-neutral-300 rounded-[6px] p-[8px]">
                Woman
              </div>
              <div class="px-[16px] hover:bg-neutral-300 rounded-[6px] p-[8px]">
                Designers
              </div>
              <div class="px-[16px] hover:bg-neutral-300 rounded-[6px] p-[8px]">
                Styles
              </div>
              <div class="px-[16px] hover:bg-neutral-300 rounded-[6px] p-[8px]">
                Trends
              </div>
            </div>
            <div class="p-4 h-[54px] w-[54px] bg-white border-slate-500 border border-opacity-50 rounded-[6px]"></div>
          </div>
          <div class="flex gap-2 items-center h-full">
            <div class="bg-neutral-300 rounded-[6px] p-0.5">
              <div class="bg-white items-center flex justify-center rounded-[6px] w-[150px] p-3 border-slate-500 border border-opacity-50 h-full font-semibold text-black text-opacity-70">
                Shop
              </div>
            </div>
            <div class="p-4 h-[54px] w-[54px] bg-white border-slate-500 border border-opacity-50 rounded-[6px]"></div>
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
                setShowInfoCard(false)
              }}
            ></div>
            <div
              id="InfoCard"
              class=" relative w-[850px] z-20 h-full p-2 gap-2 flex bg-neutral-200"
              style={{ "border-radius": "6px 0 0 6px" }}
            >
              <div
                onClick={() => {
                  setShowInfoCard(false)
                }}
                class="absolute top-3 right-3 hover:bg-neutral-400 hover:bg-opacity-40 transition-all h-[50px] w-[50px] rounded-[6px] bg-gray-200"
              ></div>
              <div class="w-[50%] h-full">
                <img src="./shoe.jpg" alt="" />
              </div>
              <div class="w-[50%] justify-between gap-8 h-full bg-white rounded-[6px] flex flex-col px-10 py-5">
                <div class="py-8">
                  <div class="flex flex-col gap-3">
                    <div class="text-[18px] opacity-90">SHOE</div>
                    <div class="text-[30px] font-light font-mono">
                      XT-Wings 2 White / White / Silver
                    </div>
                    <div class="flex gap-2 items-center text-white">
                      <div class="rounded-[4px] px-2 p-[0.5px] bg-blue-700 text-[12px] w-fit">
                        MENS
                      </div>
                      <div class="rounded-[4px] px-2 p-[0.5px] bg-blue-700 text-[12px] w-fit">
                        WOMENS
                      </div>
                      <div class="rounded-[4px] px-2 p-[0.5px] bg-blue-700 text-[12px] w-fit">
                        #LIMITED
                      </div>
                    </div>
                  </div>
                  <div class="font-light flex flex-col gap-1">
                    <div class="text-[12px]">From</div>
                    <div class="text-[20px] font-mono">$130</div>
                    <div>
                      Lorem, ipsum dolor sit amet consectetur adipisicing elit.
                      Sit nam odio perferendis molestias expedita voluptates
                      blanditiis vero nobis? Omnis voluptatem ipsum ipsam
                      voluptates nihil cum deserunt ratione unde ab temporibus.
                    </div>
                  </div>
                </div>
                <div class="w-full p-3 rounded-[6px] flex items-center justify-center font-light text-white text-[18px] bg-black">
                  Add to Cart
                </div>
              </div>
            </div>
          </div>
        </Show>
      </div>
    </>
  )
}
