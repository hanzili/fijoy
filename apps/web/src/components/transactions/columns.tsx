import { SelectTransaction } from "@/types/transaction";
import { ColumnDef } from "@tanstack/react-table";

export const columns: ColumnDef<SelectTransaction>[] = [
  {
    accessorKey: "Amount",
    header: "Amount",
  },
  {
    accessorKey: "TransactionType",
    header: "Type",
  },
  {
    accessorKey: "CategoryDetail",
    header: "Category",
    cell: ({ row }) => {
      return <div>{row.original.CategoryDetail.Name}</div>;
    },
  },
];