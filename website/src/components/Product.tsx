export default function Product() {
  return (
    <>
      <div
        id="FrontPage"
        class="w-[85%] flex flex-col bg-white dark:bg-neutral-800 h-full"
      >
        <div class="h-[100px] px-[50px] justify-between flex items-center w-full border-slate-400 border-b border-opacity-30">
          <div class="font-light text-[30px] opacity-70">Products</div>
          <div class="bg-green-500 hover:scale-[1.1] rounded-[4px] font-light text-[16px] px-[14px] p-[8px]">
            New Product
          </div>
        </div>
        <div class=" w-full flex flex-col gap-[35px] p-[50px]">
          <div class="flex items-center gap-4">
            <div class="rounded-[4px] bg-white dark:bg-black border-slate-400 border-opacity-30 border w-[250px] p-5 flex flex-col gap-1">
              <div class="font-light opacity-60 text-[14px]">Total Revenue</div>
              <div class="text-[32px] opacity-80 font-light">$250</div>
            </div>
            <div class="rounded-[4px] bg-white dark:bg-black border-slate-400 border-opacity-30 border w-[250px] p-5 flex flex-col gap-1">
              <div class="font-light opacity-60 text-[14px]">Customers</div>
              <div class="text-[32px] opacity-80 font-light">129</div>
            </div>
            <div class="rounded-[4px] bg-white dark:bg-black border-slate-400 border-opacity-30 border w-[250px] p-5 flex flex-col gap-1">
              <div class="font-light opacity-60 text-[14px]">
                Active Members
              </div>
              <div class="text-[32px] opacity-80 font-light">0</div>
            </div>
            <div class="rounded-[4px] bg-white dark:bg-black border-slate-400 border-opacity-30 border w-[250px] p-5 flex flex-col gap-1">
              <div class="font-light opacity-60 text-[14px]">MRR</div>
              <div class="text-[32px] opacity-80 font-light">$0</div>
            </div>
          </div>
        </div>
      </div>
    </>
  )
}

{
  /* <img src="solana-phone.jpg" alt="Solana phone" /> */
}
