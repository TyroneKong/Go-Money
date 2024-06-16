"use client";

import React, { useState } from "react";

import { useGetTransactions } from "@/app/hooks/requests";

import { isWithinInterval } from "date-fns";

import { SubmitHandler, useForm } from "react-hook-form";

import FormSelect from "@/ui/atoms/form-select";
import FormInput from "@/ui/atoms/form-input";

import FormFilter from "@/ui/atoms/form-filter";
import { Button } from "@/components/ui/button";
import { Download } from "lucide-react";

type SelectComponentProps = {
  columnFilter: (id: string, value: string | string[]) => void;
  id: string;
  clearFilters?: () => void;
};

function SelectComponent({ columnFilter, clearFilters }: SelectComponentProps) {
  const [open, setOpen] = useState(false);

  const { data } = useGetTransactions();

  const form = useForm<{ name: string; startDate: string; endDate: string }>({
    defaultValues: {
      name: "",
      startDate: "",
      endDate: "",
    },
  });

  const resetFields = () => {
    form.setValue("name", "");
    form.setValue("startDate", "");
    form.setValue("endDate", "");
  };

  const resetFilters = () => {
    resetFields();
    form.reset();

    clearFilters?.();
  };

  // returns an array of dates that match
  const filterDates = (startDate: string, endDate: string) => {
    const filterDate = data?.filter((x) =>
      isWithinInterval(new Date(x.createdate), {
        start: startDate,
        end: endDate,
      })
    );

    return filterDate;
  };

  type FieldTypes = {
    name: string;
    startDate: string;
    endDate: string;
  };

  const onSubmitHandler: SubmitHandler<FieldTypes> = (inputs) => {
    const matchedDates = filterDates(inputs.startDate, inputs.endDate)?.map(
      (x) => x.createdate
    );

    if (form.formState.dirtyFields.name) {
      columnFilter("name", inputs.name);
    }
    if (
      form.formState.dirtyFields.startDate &&
      form.formState.dirtyFields.endDate &&
      matchedDates
    ) {
      columnFilter("createdate", matchedDates);
    }

    setOpen(!open);
  };

  const types = Array.from(new Set(data?.map((x) => x.name))).map((name) => ({
    name,
    value: name,
  }));

  const createAndDownload = () => {
    // Convert object to a CSV-compatible array
    // const headers = Object.keys(data).join(","); // Generate CSV headers
    // const values = Object.values(data).join(","); // Generate CSV values
    // const csvContent = `${headers}\n${values}`;
    // const blob = new Blob([csvContent], {
    //   type: "text/csv;charset=utf-8",
    // });
    // const url = URL.createObjectURL(blob);
    // // Create an anchor element to download the file
    // const link = document.createElement("a");
    // link.href = url;
    // link.download = "data.csv"; // Default filename
    // link.click();
    // // Clean up the object URL
    // URL.revokeObjectURL(url);
  };

  return (
    <>
      <Button className="w-10 mb-2" onClick={createAndDownload}>
        <Download />
      </Button>
      <FormFilter
        form={form}
        onSubmit={onSubmitHandler}
        resetFilters={resetFilters}
        name="Filter"
      >
        <div className="flex flex-col gap-4">
          <FormSelect label="name" name="name" types={types} />
          <div className=" flex space-x-2">
            <div className="flex flex-col ">
              <FormInput
                label="Date Range Start"
                name="startDate"
                type="date"
              />
            </div>
            <div className="flex flex-col ">
              <FormInput label="Date Range End" name="endDate" type="date" />
            </div>
          </div>
        </div>
      </FormFilter>
    </>
  );
}

export default SelectComponent;
