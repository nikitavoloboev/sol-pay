import Image from "next/image"

export default function Card(props: { imgSrc: string; setShowInfoCard: any }) {
  return (
    <>
      <div
        className="bg-[#F0EEED] flex flex-col justify-between rounded-[6px] h-[400px] p-4"
        onClick={() => {
          props.setShowInfoCard("Info")
        }}
      >
        <div className="w-full flex items-center justify-center">
          <img src={props.imgSrc} alt="" className="max-h-[250px] bg-contain" />
        </div>
        <div className="justify-self-end">
          <div className="w-full p-2 flex flex-col items-center justify-center">
            <div className="text-[20px] font-mono">SHOE</div>
            <div className="text-[16px] font-medium opacity-50">Snake skin</div>
          </div>
          <div className="py-1 flex justify-between items-center">
            <div>Color</div>
            <div className="font-semibold text-[14px]">$710</div>
          </div>
        </div>
      </div>
    </>
  )
}
