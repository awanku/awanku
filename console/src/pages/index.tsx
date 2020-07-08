import { Box, Flex } from "@chakra-ui/core";
import { LoginForm } from "login/LoginForm";

const IndexPage = () => (
  <div>
    <Flex>
      <Box w="100%" maxW={300} p={8}>
        <LoginForm />
      </Box>
    </Flex>
  </div>
);

export default IndexPage;
