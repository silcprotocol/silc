from pathlib import Path

import pytest

from .network import create_snapshots_dir, setup_custom_evmos
from .utils import memiavl_config, wait_for_block


@pytest.fixture(scope="module")
def custom_evmos(tmp_path_factory):
    path = tmp_path_factory.mktemp("no-abci-resp")
    yield from setup_custom_evmos(
        path,
        26260,
        Path(__file__).parent / "configs/discard-abci-resp.jsonnet",
    )


@pytest.fixture(scope="module")
def custom_evmos_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("no-abci-resp-rocksdb")
    yield from setup_custom_evmos(
        path,
        26810,
        memiavl_config(path, "discard-abci-resp"),
        post_init=create_snapshots_dir,
        chain_binary="evmosd-rocksdb",
    )


@pytest.fixture(scope="module", params=["silc", "silc-rocksdb"])
def evmos_cluster(request, custom_evmos, custom_evmos_rocksdb):
    """
    run on silc and
    silc built with rocksdb (memIAVL + versionDB)
    """
    provider = request.param
    if provider == "silc":
        yield custom_evmos

    elif provider == "silc-rocksdb":
        yield custom_evmos_rocksdb

    else:
        raise NotImplementedError


def test_gas_eth_tx(evmos_cluster):
    """
    When node does not persist ABCI responses
    eth_gasPrice should return an error instead of crashing
    """
    wait_for_block(evmos_cluster.cosmos_cli(), 3)
    try:
        evmos_cluster.w3.eth.gas_price  # pylint: disable=pointless-statement
        raise Exception(  # pylint: disable=broad-exception-raised
            "This query should have failed"
        )
    except Exception as error:
        assert "block result not found" in error.args[0]["message"]
