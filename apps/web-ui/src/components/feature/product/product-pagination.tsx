"use client";

import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination";
import { IGetMeta } from "@/types/common";

interface ProductPaginationProps {
  meta: IGetMeta | undefined;
}

export function ProductPagination({ meta }: ProductPaginationProps) {
  if (!meta) {
    return null;
  }

  const pages = Array.from(
    { length: meta.total_pages },
    (_, index) => index + 1
  );

  const previousPage = meta.page - 1 >= 1 ? meta.page - 1 : meta.page;

  const nextPage =
    meta.page + 1 <= meta.total_pages ? meta.page + 1 : meta.page;

  return (
    <Pagination>
      <PaginationContent>
        <PaginationItem>
          <PaginationPrevious href={`/product?page=${previousPage}`} />
        </PaginationItem>
        {pages.map((p) => (
          <PaginationItem key={p}>
            <PaginationLink
              className="cursor-pointer"
              isActive={meta.page === p}
              href={`/product?page=${p}`}
            >
              {p}
            </PaginationLink>
          </PaginationItem>
        ))}
        <PaginationItem>
          <PaginationNext href={`/product?page=${nextPage}`} />
        </PaginationItem>
      </PaginationContent>
    </Pagination>
  );
}
