import Nav from "~/components/Nav"
import Product from "~/components/Product"

export default function Home() {
  return (
    <main class="w-screen h-full flex flex-col gap-16">
      <Nav />
      <div class="px-[10%] w-full h-full">
        <div class="border-2 w-full h-full rounded-[4px] border-black flex flex-col">
          <div
            style={{ "background-image": "../public/solana-phone.jpg" }}
            class="px-[50px] border-b-2  border-black p-[50px] h-[500px] text-[46px] font-bold flex flex-col justify-center items-center w-full"
          >
            <div>
              <div>Buy Phone Please!</div>
              <div class="text-[#696969] font-light text-[32px]">
                Its got a great camera
              </div>
            </div>
          </div>

          <div class="w-full h-full relative">
            <div
              id="StickyBar"
              class="p-6 border-b-2 bg-white z-10 sticky top-0 left-0 border-black flex justify-between items-center"
            >
              <div class="text-[20px] flex items-center gap-3">
                <div class="bg-black rounded-[4px] text-[16px] px-2 p-0.5 text-white">
                  $299
                </div>
                <div>Phone seller</div>
              </div>
              <div class="flex items-center gap-3">
                <div>stars</div>
                <div class="px-[14px] p-[8px] font-light bg-black rounded-[4px] text-white">
                  Add to cart
                </div>
              </div>
            </div>
            <div class="w-full h-full flex">
              <div class="w-[60%] h-[1000px] border-r-2 border-black">
                <div>TYPE HERE</div>
              </div>
              <div class="w-[40%] flex flex-col">
                <div id="Value" class="w-full p-6 border-b-2 border-black">
                  <div class="rounded-[4px] border-2 border-black flex flex-col font-light text-[18px]">
                    <div class="border-b-2 border-black px-5 p-3 flex items-center justify-between">
                      You'll get:
                    </div>
                    <div class="border-b-2 border-black px-5 p-3 flex items-center justify-between">
                      <div>Access camera</div>
                      <div class="opacity-60 text-[16px]">Instant</div>
                    </div>
                    <div class="border-b-2 border-black px-5 p-3 flex items-center justify-between">
                      <div>New Phone</div>
                      <div class="opacity-60 text-[16px]">Instant</div>
                    </div>
                    <div class="px-5 p-3 flex items-center justify-between">
                      <div>Phone Case</div>
                      <div class="opacity-60 text-[16px]">Included</div>
                    </div>
                  </div>
                </div>
                <div
                  id="Rating"
                  class="flex gap-4 flex-col w-full h-full p-6 px-8 border-b-2 border-black"
                >
                  <div class="text-[20px] flex items-center justify-between">
                    <div>Ratings</div>
                    <div class="text-[16px] opacity-60">2.1 ratings</div>
                  </div>
                  <div
                    id="RatingBar"
                    class="w-full flex items-center justify-between gap-2"
                  >
                    <div class="flex gap-1">
                      <div>5</div>
                      <div>stars</div>
                    </div>
                    <div class="w-full h-full rounded-[4px] bg-black border-black border-2"></div>
                    <div class="w-[70px] flex items-center justify-center">
                      100%
                    </div>
                  </div>
                  <div
                    id="RatingBar"
                    class="w-full flex items-center justify-between gap-2"
                  >
                    <div class="flex gap-1">
                      <div>4</div>
                      <div>stars</div>
                    </div>
                    <div class="w-full h-full rounded-[4px] border-black border-2"></div>
                    <div class="w-[70px] flex items-center justify-center">
                      0%
                    </div>
                  </div>
                  <div
                    id="RatingBar"
                    class="w-full flex items-center justify-between gap-2"
                  >
                    <div class="flex gap-1">
                      <div>3</div>
                      <div>stars</div>
                    </div>
                    <div class="w-full h-full rounded-[4px]  border-black border-2"></div>
                    <div class="w-[70px] flex items-center justify-center">
                      0%
                    </div>
                  </div>
                  <div
                    id="RatingBar"
                    class="w-full flex items-center justify-between gap-2"
                  >
                    <div class="flex gap-1">
                      <div>2</div>
                      <div>stars</div>
                    </div>
                    <div class="w-full h-full rounded-[4px]  border-black border-2"></div>
                    <div class="w-[70px] flex items-center justify-center">
                      0%
                    </div>
                  </div>
                  <div
                    id="RatingBar"
                    class="w-full flex items-center justify-between gap-2"
                  >
                    <div class="flex gap-1">
                      <div>1</div>
                      <div>stars</div>
                    </div>
                    <div class="w-full h-full rounded-[4px]  border-black border-2"></div>
                    <div class="w-[70px] flex items-center justify-center">
                      0%
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  )
}
