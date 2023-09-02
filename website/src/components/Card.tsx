export default function Card(props: any) {
  return (
    <>
      <div
        class="bg-[#F0EEED] flex flex-col justify-between rounded-[6px] h-[400px] p-4"
        onClick={() => {
          props.setShowInfoCard("Info")
        }}
      >
        <div class="w-full flex items-center justify-center">
          <img src="./shoe.jpg" alt="" class="max-h-[250px] bg-contain" />
        </div>
        <div class="justify-self-end">
          <div class="w-full p-2 flex flex-col items-center justify-center">
            <div class="text-[20px] font-mono">SHOE</div>
            <div class="text-[16px] font-medium opacity-50">Snake skin</div>
          </div>
          <div class="py-1 flex justify-between items-center">
            <div>Color</div>
            <div class="font-semibold text-[14px]">$710</div>
          </div>
        </div>
      </div>
    </>
  )
}
