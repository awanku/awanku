import {
  Box,
  Button,
  FormControl,
  FormHelperText,
  FormLabel,
  Input,
} from "@chakra-ui/core";

export let LoginForm = () => {
  return (
    <Box>
      <a href="https://api.awanku.id/v1/auth/google/connect?redirect_to=http://localhost:3000&state=google">
        <Button>Continue with Google</Button>
      </a>
      <a href="https://api.awanku.id/v1/auth/github/connect?redirect_to=http://localhost:3000&state=github">
        <Button>Continue with Github</Button>
      </a>
      <FormControl>
        <FormLabel htmlFor="email">Email address</FormLabel>
        <Input type="email" id="email" aria-describedby="email-helper-text" />
        <FormHelperText id="email-helper-text">
          We'll never share your email.
        </FormHelperText>
      </FormControl>
    </Box>
  );
};
