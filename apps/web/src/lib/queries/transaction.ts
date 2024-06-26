import { Transport } from "@connectrpc/connect";
import { createQueryOptions } from "@connectrpc/connect-query";
import { Workspace } from "@/gen/proto/fijoy/v1/workspace_pb";
import { getTransactions } from "@/gen/proto/fijoy/v1/transaction-TransactionService_connectquery";
import { getWorkspaceHeader } from "@/lib/headers";

type getTransactionsProps = {
  context: {
    transport: Transport;
    workspace: Workspace;
  };
};

export const getTransactionsQueryOptions = ({
  context,
}: getTransactionsProps) => {
  return createQueryOptions(
    getTransactions,
    {},
    {
      transport: context.transport,
      callOptions: {
        headers: getWorkspaceHeader(context.workspace.id),
      },
    },
  );
};
