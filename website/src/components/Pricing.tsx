export default function Pricing(props: any) {
  return (
    <div class="w-full flex">
      {/* <div
        onClick={() => {
          setShowInfoCard(false)
        }}
        class="absolute top-3 right-3 hover:bg-neutral-400 hover:bg-opacity-40 transition-all h-[50px] w-[50px] rounded-[6px] bg-gray-200"
      ></div> */}
      <div class="w-[50%] h-full flex flex-col gap-6 items-center py-14 px-6">
        <div class="text-[30px] font-semibold self-start">Title</div>
        <div class="flex flex-col">
          <div class="w-full text-[14px] font-light opacity-90 flex justify-between items-center">
            <div>Software</div>
            <div>stars</div>
          </div>
          <div class="flex flex-col gap-2">
            <div class="text-[30px] font-mono">Description</div>
            <div class="font-light text-[16px] ">
              Lorem ipsum dolor sit amet consectetur adipisicing elit. Deserunt
              veritatis rem cupiditate dicta nobis, quisquam, inventore optio
              neque provident aperiam quas, a quidem nisi? Quas voluptatum cum
              asperiores neque dicta!
            </div>
          </div>
        </div>
        <div class="w-full flex justify-between items-center">
          <div>From 200$</div>
          <div>Name of seller</div>
        </div>
      </div>
      <div class="w-[50%] justify-between gap-8 h-full bg-white rounded-[6px] flex flex-col px-10 py-5">
        <div class="py-8 flex flex-col gap-10">
          <div class="flex flex-col gap-2">
            <div class="font-light text-[26px]">
              Current balance: <span class="font-semibold">$30</span>
            </div>
            <div class="flex font-light flex-col gap-2">
              <div class="text-[16px]">SOL address: rtogkgr</div>
              <div class="text-[16px]">
                Please top up your account with{" "}
                <span class="font-semibold">$170</span> to continue
              </div>
            </div>
          </div>
          <div class="flex flex-col gap-2">
            <div class="flex gap-2 items-center">
              <div class="font-light text-[18px]">Price</div>
              <input
                type="text"
                class="w-3/5 text-opacity-70 text-black outline-none px-1 p-1 rounded-[4px] border border-slate-400 border-opacity-50"
              />
            </div>
            <div class="text-[12px] opacity-70">Write price from $200</div>
          </div>
        </div>
        <div
          onClick={() => {
            props.setShowInfoCard("Bought")
          }}
          class="w-full p-3 rounded-[6px] flex items-center justify-center font-light text-white text-[18px] bg-black"
        >
          Pay
        </div>
      </div>
    </div>
  )
}
