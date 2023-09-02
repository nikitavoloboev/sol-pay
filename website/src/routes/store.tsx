import { Show, createSignal, onMount } from "solid-js"
import Card from "~/components/Card"
import { Core } from "@walletconnect/core"
import { Web3Wallet } from "@walletconnect/web3wallet"
import { json } from "solid-start"

export default function store() {
  const [showInfoCard, setShowInfoCard] = createSignal(true)
  const [notEnoughMoney, setNotEnoughMoney] = createSignal(false)
  const [checkedBalance, setCheckedBalance] = createSignal(false)

  // onMount(async () => {
  //   const url = "http://localhost:3000/balance"
  //   const data = {
  //     buyer_id: 8,
  //     product_id: 2,
  //   }

  //   const response = await fetch(url, {
  //     method: "POST",
  //     headers: {
  //       "Content-Type": "application/json",
  //     },
  //     body: JSON.stringify(data),
  //   })

  //   if (response.ok) {
  //     const jsonResponse = await response.json()
  //     console.log("Payment successful:", jsonResponse)
  //   } else {
  //     console.log("Payment failed:", response.status, response.statusText)
  //   }
  // })
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
                {/* <div
                  onClick={async () => {
                    // const core = new Core({
                    //   projectId: ,
                    // })
                    // const web3wallet = await Web3Wallet.init({
                    //   core, // <- pass the shared `core` instance
                    //   metadata: {
                    //     name: "Demo app",
                    //     description: "Demo Client as Wallet/Peer",
                    //     url: "www.walletconnect.com",
                    //     icons: [],
                    //   },
                    // })
                  }}
                >
                  Connect wallet
                </div> */}

                <div
                  onClick={async () => {
                    if (checkedBalance()) {
                      const url = "http://localhost:8080/pay"
                      const data = {
                        buyer_id: 8,
                        seller_id: 2,
                        product_id: 3,
                      }

                      const response = await fetch(url, {
                        method: "POST",
                        headers: {
                          "Content-Type": "application/json",
                        },
                        body: JSON.stringify(data),
                      })

                      if (response.ok) {
                        const jsonResponse = await response.json()
                        console.log("Payment successful:", jsonResponse)
                      } else {
                        console.log(
                          "Payment failed:",
                          response.status,
                          response.statusText
                        )
                      }
                    } else {
                      const url = "http://localhost:8080/can-i-buy"
                      /* DK FIXME - TO BE CHANGED TO IDs of local DB! */
                      const data = {
                        user_id: 20,
                        product_id: 9,
                      }

                      const response = await fetch(url, {
                        method: "POST",
                        headers: {
                          "Content-Type": "application/json",
                          "Access-Control-Allow-Origin": "*",
                        },
                        body: JSON.stringify(data),
                      })
                      console.log(response, "response")
                      console.log(await response.text())

                      const jsonResponse = await response.json()
                      console.log(jsonResponse, "json response")

                      if (jsonResponse.can_pay) {
                        const jsonResponse = await response.json()
                        console.log("Can pay:", jsonResponse)
                      } else {
                        console.log(
                          "Cannot pay",
                          response.status,
                          response.statusText
                        )
                      }
                    }
                  }}
                  class="w-full p-3 rounded-[6px] flex items-center justify-center font-light text-white text-[18px] bg-black"
                >
                  Pay
                </div>
              </div>
            </div>
          </div>
        </Show>
      </div>
    </>
  )
}
