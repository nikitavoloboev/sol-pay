export default function FrontPage() {
  return (
    <>
      <div
        id="FrontPage"
        class="w-[85%] flex flex-col bg-white dark:bg-neutral-800 h-full"
      >
        <div class="h-[100px] px-[50px] justify-between flex items-center w-full border-slate-400 border-b border-opacity-30">
          <div class="font-light text-[30px] opacity-70">Title</div>
          <div class="bg-green-500 hover:scale-[1.1] rounded-[4px] font-light text-[16px] px-[14px] p-[8px]">
            Button
          </div>
        </div>
        <div class="w-full flex p-[50px] gap-[50px]">
          <div class="w-[20%] font-light text-[14px] opacity-60">
            Lorem ipsum dolor sit, amet consectetur adipisicing elit.
            Reprehenderit molestias iure suscipit ipsam incidunt excepturi
            laborum maxime libero, recusandae iusto, reiciendis repellat
            voluptatibus sunt esse quam. Pariatur impedit aliquam est!
          </div>
          <div class="w-[80%] flex flex-col gap-[35px]">
            <div class="w-full flex flex-col gap-2">
              <div class="font-light text-[16px] opacity-60">Name</div>
              <input
                type="text"
                class="w-full rounded-[4px] bg-gray-200 dark:bg-black border-slate-400 border-opacity-30 border p-[10px] px-[12px] outline-none text-opacity-50 text-black dark:text-white text-[14px]"
              />
            </div>
            <div class="flex flex-col gap-2">
              <div class="font-light text-[16px] opacity-60">Type</div>
              <div class="flex gap-4">
                <div class="p-4 rounded-[4px]  border-slate-400 border border-opacity-30 w-fit">
                  <div class="font-light text-[16px] opacity-60">
                    Digital product
                  </div>
                  <div class="font-light text-[14px] opacity-40">
                    Any set of files to download or stream
                  </div>
                </div>
                <div class="p-4 rounded-[4px]  border-slate-400 border border-opacity-30 w-fit">
                  <div class="font-light text-[16px] opacity-60">Ideas</div>
                  <div class="font-light text-[14px] opacity-40">idfk</div>
                </div>
              </div>
            </div>
            <div class="flex flex-col gap-2">
              <div class="font-light text-[16px] opacity-60">Price</div>
              <input
                type="text"
                class="w-full rounded-[4px] bg-gray-200 dark:bg-black border-slate-400 border-opacity-30 border p-[10px] px-[12px] outline-none text-opacity-50 text-black dark:text-white text-[14px]"
              />
            </div>
          </div>
        </div>
      </div>
    </>
  )
}
