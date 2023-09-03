export default function Info(props: any) {
  return (
    <div className="w-full flex">
      {/* <div
        onClick={() => {
          setShowInfoCard(false)
        }}
        className="absolute top-3 right-3 hover:bg-neutral-400 hover:bg-opacity-40 transition-all h-[50px] w-[50px] rounded-[6px] bg-gray-200"
      ></div> */}
      <div className="w-[50%] h-full flex flex-col gap-6 items-center py-14 px-6">
        <div className="text-[30px] font-semibold self-start">Title</div>
        <div className="w-full gap-2 flex flex-col">
          <div className="w-full text-[14px] font-light opacity-90 flex justify-between items-center">
            <div>Software</div>
            <div>stars</div>
          </div>
          <div className="w-full h-[300px] rounded-[6px] border border-slate-400 border-opacity-50"></div>
        </div>
        <div className="w-full flex justify-between items-center">
          <div>From 200$</div>
          <div>Name of seller</div>
        </div>
      </div>
      <div className="w-[50%] justify-between gap-8 h-full bg-white rounded-[6px] flex flex-col px-10 py-5">
        <div className="flex flex-col gap-2 py-8">
          <div className="text-[30px] font-mono">Description</div>
          <div className="font-light text-[16px] ">
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Deserunt
            veritatis rem cupiditate dicta nobis, quisquam, inventore optio
            neque provident aperiam quas, a quidem nisi? Quas voluptatum cum
            asperiores neque dicta!
          </div>
        </div>
        <div
          onClick={() => {
            props.setShowInfoCard("Pricing")
          }}
          className="w-full p-3 rounded-[6px] flex items-center justify-center font-light text-white text-[18px] bg-black"
        >
          Pay
        </div>
      </div>
    </div>
  )
}
