import { AppProps } from "next/app"
import Head from "next/head"
import { FC } from "react"
import { ContextProvider } from "../contexts/ContextProvider"
import { AppBar } from "../components/AppBar"
import { ContentContainer } from "../components/ContentContainer"
import { Footer } from "../components/Footer"
import Notifications from "../components/Notification"
import Store from "components/store/Store"
require("@solana/wallet-adapter-react-ui/styles.css")
require("../styles/globals.css")

const App: FC<AppProps> = ({ Component, pageProps }) => {
  return (
    <>
      <Head>
        <title>IdeaStub</title>
      </Head>

      <ContextProvider>
        <div className="flex flex-col h-screen w-screen bg-white text-black">
          <Notifications />
          <AppBar />

          <Store />
        </div>
      </ContextProvider>
    </>
  )
}

export default App
