(window.webpackJsonp=window.webpackJsonp||[]).push([[57],{410:function(t,e,a){"use strict";a.r(e);var s=a(9),n=Object(s.a)({},(function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[a("h1",{attrs:{id:"setting-up-tendermint-kms-ledger"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#setting-up-tendermint-kms-ledger"}},[t._v("#")]),t._v(" Setting up Tendermint KMS + Ledger")]),t._v(" "),a("div",{staticClass:"custom-block danger"},[a("p",{staticClass:"custom-block-title"},[t._v("Warning")]),t._v(" "),a("p",[t._v("The following instructions are a brief walkthrough and not a comprehensive guideline. You should consider and "),a("RouterLink",{attrs:{to:"/validators/security.html"}},[t._v("research more about the security implications")]),t._v(" of activating an external KMS.")],1)]),t._v(" "),a("div",{staticClass:"custom-block danger"},[a("p",{staticClass:"custom-block-title"},[t._v("Warning")]),t._v(" "),a("p",[t._v("KMS and Ledger Tendermint app are currently work in progress. Details may vary. Use with care under your own risk.")])]),t._v(" "),a("h2",{attrs:{id:"tendermint-validator-app-for-ledger-devices"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#tendermint-validator-app-for-ledger-devices"}},[t._v("#")]),t._v(" Tendermint Validator app (for Ledger devices)")]),t._v(" "),a("p",[t._v("You should be able to find the Tendermint app in Ledger Live.")]),t._v(" "),a("p",[a("em",[t._v("Note: at the moment, you might need to enable "),a("code",[t._v("developer mode")]),t._v(" in Ledger Live settings")])]),t._v(" "),a("h2",{attrs:{id:"kms-configuration"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#kms-configuration"}},[t._v("#")]),t._v(" KMS configuration")]),t._v(" "),a("p",[t._v("In this section, we will configure a KMS to use a Ledger device running the Tendermint Validator App.")]),t._v(" "),a("h3",{attrs:{id:"config-file"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#config-file"}},[t._v("#")]),t._v(" Config file")]),t._v(" "),a("p",[t._v("You can find other configuration examples "),a("a",{attrs:{href:"https://github.com/iqlusioninc/tmkms/blob/develop/tmkms.toml.example",target:"_blank",rel:"noopener noreferrer"}},[t._v("here"),a("OutboundLink")],1)]),t._v(" "),a("ul",[a("li",[a("p",[t._v("Create a "),a("code",[t._v("~/.tmkms/tmkms.toml")]),t._v(" file with the following content (use an adequate "),a("code",[t._v("chain_id")]),t._v(")")]),t._v(" "),a("div",{staticClass:"language-toml line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-toml"}},[a("code",[a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# Example KMS configuration file")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),a("span",{pre:!0,attrs:{class:"token table class-name"}},[t._v("validator")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token key property"}},[t._v("addr")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"tcp://localhost:26658"')]),t._v("    "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v('# or "unix:///path/to/socket"')]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token key property"}},[t._v("chain_id")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"gaia-11001"')]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token key property"}},[t._v("reconnect")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token boolean"}},[t._v("true")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# true is the default")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token key property"}},[t._v("secret_key")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"~/.tmkms/secret_connection.key"')]),t._v("\n\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),a("span",{pre:!0,attrs:{class:"token table class-name"}},[t._v("providers.ledgertm")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token key property"}},[t._v("chain_ids")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("[")]),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"gaia-11001"')]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("]")]),t._v("\n")])]),t._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[t._v("1")]),a("br"),a("span",{staticClass:"line-number"},[t._v("2")]),a("br"),a("span",{staticClass:"line-number"},[t._v("3")]),a("br"),a("span",{staticClass:"line-number"},[t._v("4")]),a("br"),a("span",{staticClass:"line-number"},[t._v("5")]),a("br"),a("span",{staticClass:"line-number"},[t._v("6")]),a("br"),a("span",{staticClass:"line-number"},[t._v("7")]),a("br"),a("span",{staticClass:"line-number"},[t._v("8")]),a("br"),a("span",{staticClass:"line-number"},[t._v("9")]),a("br")])])]),t._v(" "),a("li",[a("p",[t._v("Edit "),a("code",[t._v("addr")]),t._v(" to point to your "),a("code",[t._v("gaiad")]),t._v(" instance.")])]),t._v(" "),a("li",[a("p",[t._v("Adjust "),a("code",[t._v("chain-id")]),t._v(" to match your "),a("code",[t._v(".gaiad/config/config.toml")]),t._v(" settings.")])]),t._v(" "),a("li",[a("p",[a("code",[t._v("provider.ledgertm")]),t._v(" has not additional parameters at the moment, however, it is important that you keep that header to enable the feature.")])])]),t._v(" "),a("p",[a("em",[t._v("Plug your Ledger device and open the Tendermint validator app.")])]),t._v(" "),a("h3",{attrs:{id:"generate-secret-key"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#generate-secret-key"}},[t._v("#")]),t._v(" Generate secret key")]),t._v(" "),a("p",[t._v("Now you need to generate secret_key:")]),t._v(" "),a("div",{staticClass:"language-bash line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[t._v("tmkms keygen ~/.tmkms/secret_connection.key\n")])]),t._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[t._v("1")]),a("br")])]),a("h3",{attrs:{id:"retrieve-validator-key"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#retrieve-validator-key"}},[t._v("#")]),t._v(" Retrieve validator key")]),t._v(" "),a("p",[t._v("The last step is to retrieve the validator key that you will use in "),a("code",[t._v("gaiad")]),t._v(".")]),t._v(" "),a("p",[t._v("Start the KMS:")]),t._v(" "),a("div",{staticClass:"language-bash line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[t._v("tmkms start -c ~/.tmkms/tmkms.toml\n")])]),t._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[t._v("1")]),a("br")])]),a("p",[t._v("The output should look similar to:")]),t._v(" "),a("div",{staticClass:"language- line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[t._v("07:28:24 [INFO] tmkms 0.3.0 starting up...\n07:28:24 [INFO] [keyring:ledgertm:ledgertm] added validator key cosmosvalconspub1zcjduepqy53m39prgp9dz3nz96kaav3el5e0th8ltwcf8cpavqdvpxgr5slsd6wz6f\n07:28:24 [INFO] KMS node ID: 1BC12314E2E1C29015B66017A397F170C6ECDE4A\n")])]),t._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[t._v("1")]),a("br"),a("span",{staticClass:"line-number"},[t._v("2")]),a("br"),a("span",{staticClass:"line-number"},[t._v("3")]),a("br")])]),a("p",[t._v("The KMS may complain that it cannot connect to gaiad. That is fine, we will fix it in the next section.")]),t._v(" "),a("p",[t._v("This output indicates the validator key linked to this particular device is: "),a("code",[t._v("cosmosvalconspub1zcjduepqy53m39prgp9dz3nz96kaav3el5e0th8ltwcf8cpavqdvpxgr5slsd6wz6f")])]),t._v(" "),a("p",[t._v("Take note of the validator pubkey that appears in your screen. "),a("em",[t._v("We will use it in the next section.")])]),t._v(" "),a("h2",{attrs:{id:"gaia-configuration"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#gaia-configuration"}},[t._v("#")]),t._v(" Gaia configuration")]),t._v(" "),a("p",[t._v("You need to enable KMS access by editing "),a("code",[t._v(".gaiad/config/config.toml")]),t._v(". In this file, modify "),a("code",[t._v("priv_validator_laddr")]),t._v(" to create a listening address/port or a unix socket in "),a("code",[t._v("gaiad")]),t._v(".")]),t._v(" "),a("p",[t._v("For example:")]),t._v(" "),a("div",{staticClass:"language-toml line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-toml"}},[a("code",[a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# TCP or UNIX socket address for Tendermint to listen on for")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token comment"}},[t._v("# connections from an external PrivValidator process")]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token key property"}},[t._v("priv_validator_laddr")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("=")]),t._v(" "),a("span",{pre:!0,attrs:{class:"token string"}},[t._v('"tcp://127.0.0.1:26658"')]),t._v("\n"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v(".")]),t._v("\n")])]),t._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[t._v("1")]),a("br"),a("span",{staticClass:"line-number"},[t._v("2")]),a("br"),a("span",{staticClass:"line-number"},[t._v("3")]),a("br"),a("span",{staticClass:"line-number"},[t._v("4")]),a("br"),a("span",{staticClass:"line-number"},[t._v("5")]),a("br")])]),a("p",[t._v("Let's assume that you have set up your validator account and called it "),a("code",[t._v("kmsval")]),t._v(". You can tell gaiad the key that we've got in the previous section.")]),t._v(" "),a("div",{staticClass:"language-bash line-numbers-mode"},[a("pre",{pre:!0,attrs:{class:"language-bash"}},[a("code",[t._v("gaiad gentx --name kmsval --pubkey "),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("{")]),t._v(".ValidatorKey"),a("span",{pre:!0,attrs:{class:"token punctuation"}},[t._v("}")]),t._v(" \n")])]),t._v(" "),a("div",{staticClass:"line-numbers-wrapper"},[a("span",{staticClass:"line-number"},[t._v("1")]),a("br")])]),a("p",[t._v("Now start "),a("code",[t._v("gaiad")]),t._v(". You should see that the KMS connects and receives a signature request.")]),t._v(" "),a("p",[t._v("Once the ledger receives the first message, it will ask for confirmation that the values are adequate.")]),t._v(" "),a("p",[a("img",{attrs:{src:"ledger_1.jpg",alt:""}})]),t._v(" "),a("p",[t._v("Click the right button, if the height and round are correct.")]),t._v(" "),a("p",[t._v("After that, you will see that the KMS will start forwarding all signature requests to the ledger:")]),t._v(" "),a("p",[a("img",{attrs:{src:"ledger_2.jpg",alt:""}})]),t._v(" "),a("div",{staticClass:"custom-block danger"},[a("p",{staticClass:"custom-block-title"},[t._v("Warning")]),t._v(" "),a("p",[t._v("The word TEST in the second picture, second line appears because they were taken on a pre-release version.")]),t._v(" "),a("p",[t._v("Once the app as been released in Ledger's app store, this word should NOT appear.")])])])}),[],!1,null,null,null);e.default=n.exports}}]);