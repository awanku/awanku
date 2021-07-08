const fetch = require("node-fetch");
const codegen = require("swagger-typescript-codegen");
const fs = require("fs");
const path = require("path");

fetch("https://api.awanku.id/docs/swagger.json")
  .then((res) => res.json())
  .then((swagger) => {
    const tsSourceCode = codegen.CodeGen.getTypescriptCode({
      className: "API",
      swagger,
    });

    let content = tsSourceCode.replace(
      /export type (([a-zA-Z0-9]+)\.([a-zA-Z0-9]+))/g,
      "export type $2_$3"
    );

    let content2 = content.replace(
      /((hansip|core|coreapi|auth|oauth2|apihelper)\.([a-zA-Z0-9]+))/g,
      "$2_$3"
    );

    fs.writeFileSync(
      path.join(__dirname, "../src/__generated__/API.ts"),
      content2,
      {
        encoding: "utf8",
      }
    );
  })
  .catch((err) => console.error(err));
