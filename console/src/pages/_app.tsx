import { ColorModeProvider, CSSReset, ThemeProvider } from "@chakra-ui/core";
import NextApp from "next/app";

class App extends NextApp {
  render() {
    const { Component, pageProps } = this.props;
    return (
      <ThemeProvider>
        <ColorModeProvider value="light">
          <CSSReset />
          <Component {...pageProps} />
        </ColorModeProvider>
      </ThemeProvider>
    );
  }
}

export default App;
