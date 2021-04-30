#!/bin/sh
#set -o errexit -o nounset -o pipefail

KEY_PASSWORD=${KEY_PASSWORD:-1234567890}
KEY_NAME=${KEY_NAME:-validator}
STAKE=${STAKE_TOKEN:-udarc}
FEE=${FEE_TOKEN:-udarc}
CHAIN_ID=${CHAIN_ID:-darchub}
MONIKER=${MONIKER:-dm0}

# ------------------------------------------------------------------------------
#
# Print env variables
#
# ------------------------------------------------------------------------------
echo "Chain id" "${CHAIN_ID}"
echo "Moniker" "${MONIKER}"
echo "Stake denom" "${STAKE}"
echo "Fee denom" "${FEE}"
echo "Key name" "${KEY_NAME}"
echo "Mnemonic" "${KEY_MNEMONIC}"
echo "Password" "${KEY_PASSWORD}"

if test -n "${KEY_MNEMONIC-}"
then
  echo "$KEY_MNEMONIC"
    {
      echo "${KEY_MNEMONIC}"
      echo "${KEY_PASSWORD}"
      echo "${KEY_PASSWORD}"
      echo
    } | knstld keys add "${KEY_NAME}" --recover --log_level info
  # hardcode the validator account for this instance
else
    {
      echo "${KEY_PASSWORD}"
      echo "${KEY_PASSWORD}"
      echo
    } | knstld keys add "${KEY_NAME}" --log_level info
fi
echo "$KEY_PASSWORD" | knstld add-genesis-account "$KEY_NAME" "200000000000$STAKE" --log_level info

# (optionally) add a few more genesis accounts
for addr in "$@"; do
  echo $addr
  knstld add-genesis-account "$addr" "200000000000$STAKE" --log_level info
done
# submit a genesis validator tx
## Workraround for https://github.com/cosmos/cosmos-sdk/issues/8251
(echo "$KEY_PASSWORD"; echo "$KEY_PASSWORD"; echo "$KEY_PASSWORD") | knstld gentx "$KEY_NAME" "100000000000$STAKE" --chain-id="$CHAIN_ID" --amount="100000000000$STAKE" --log_level info
## should be:
# (echo "$PASSWORD"; echo "$PASSWORD"; echo "$PASSWORD") | knstld gentx validator "250000000$STAKE" --chain-id="$CHAIN_ID"
knstld collect-gentxs --log_level info
