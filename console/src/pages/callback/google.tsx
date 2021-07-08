import { Flex, Spinner } from "@chakra-ui/core";
import { api } from "api";
import { GetServerSideProps } from "next";
import React from "react";

type Props = {
  code?: string | string[];
};

let CallbackPage: React.FunctionComponent<Props> = (props) => {
  React.useEffect(() => {
    api.api_v1_auth_exchangeToken({
      param: {
        code: props.code as string,
        grant_type: "authorization_code",
        refresh_token: "",
      },
    });
  }, []);

  return (
    <Flex w="100%" h="100vh" alignItems="center" justifyContent="center">
      <Spinner />
    </Flex>
  );
};

export const getServerSideProps: GetServerSideProps<Props> = async (
  context
) => {
  let code = context.query.code;
  return {
    props: {
      code,
    },
  };
};

export default CallbackPage;
