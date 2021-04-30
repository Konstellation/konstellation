#!/bin/sh
#set -o errexit -o nounset -o pipefail

STAKE=${STAKE_TOKEN:-udarc}
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

knstld init --chain-id "$CHAIN_ID" "$MONIKER" --log_level info
# staking/governance token is hardcoded in config, change this
sed -i "s/\"stake\"/\"$STAKE\"/" "$HOME"/.knstld/config/genesis.json
