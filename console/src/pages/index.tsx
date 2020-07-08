import { Box, Button, Flex, useTheme } from "@chakra-ui/core";

let connectUrl = (type: "google" | "github") => {
  if (process.browser) {
    let redirectTo = window.location.origin + "/callback/" + type;
    window.location.replace(
      `https://api.awanku.id/v1/auth/${type}/connect?redirect_to=${redirectTo}&state=index`
    );
  }
};

const IndexPage = () => {
  let theme = useTheme();
  return (
    <Flex
      w="100%"
      h="100vh"
      alignItems="center"
      justifyContent="center"
      background={theme.colors.gray[100]}
    >
      <Box
        w="100%"
        maxW={300}
        p={8}
        shadow="sm"
        background={theme.colors.white}
      >
        <Button onClick={() => connectUrl("google")} isFullWidth>
          Continue with Google
        </Button>
        <Box h={4} />
        <Button onClick={() => connectUrl("google")} isFullWidth>
          Continue with Github
        </Button>
      </Box>
    </Flex>
  );
};

export default IndexPage;
