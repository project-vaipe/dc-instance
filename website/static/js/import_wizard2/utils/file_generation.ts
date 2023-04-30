/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
import _ from "lodash";
import Papa from "papaparse";

import { CsvData, Mapping, MappingVal, ValueMap } from "../types";

// convert Mapping object to a regular javascript object so it can be converted
// into a JSON string
function mappingToObject(mapping: Mapping): Record<string, MappingVal> {
  const obj = {};
  mapping.forEach((mappingVal, mappedThing) => {
    obj[mappedThing] = mappingVal;
  });
  return obj;
}

/**
 * Generates the translation metadata json string given the predicted mappings
 * and correct mappings
 * @param predictions
 * @param correctedMapping
 * @returns
 */
export function generateTranslationMetadataJson(
  predictions: Mapping,
  correctedMapping: Mapping
): string {
  const translationMetadata = {
    predictions: mappingToObject(predictions),
    correctedMapping: mappingToObject(correctedMapping),
  };
  return JSON.stringify(translationMetadata);
}

/**
 * Returns whether or not Csv should be generated. Csv needs to be generated if
 * any of the following is true:
 *    - the column header has been updated
 *    - if the column id was auto generated
 *    - if valueMap is not empty
 *    - if header Row isn't 1, first data row isn't 2 and last data row isn't
 *      the last row in the file
 * @param originalCsv
 * @param correctedCsv
 * @param valueMap
 * @returns
 */
export function shouldGenerateCsv(
  originalCsv: CsvData,
  correctedCsv: CsvData,
  valueMap: ValueMap
): boolean {
  if (!_.isEmpty(valueMap)) {
    return true;
  }
  // We need to generate a cleaned CSV if the column header has been updated
  // or if the column id was auto generated.
  let hasUpdatedColumnId = false;
  correctedCsv.orderedColumns.forEach((column, idx) => {
    const columnEditedByUser = column.id !== originalCsv.orderedColumns[idx].id;
    const columnIdAutoGenerated = column.id !== column.header;
    if (columnEditedByUser || columnIdAutoGenerated) {
      hasUpdatedColumnId = true;
      return;
    }
  });
  return (
    hasUpdatedColumnId ||
    correctedCsv.headerRow !== 1 ||
    correctedCsv.firstDataRow !== 2 ||
    correctedCsv.lastDataRow !== correctedCsv.lastFileRow
  );
}

/**
 * Returns a promise with a Csv Blob.
 * @param csvData
 */
export function generateCsv(
  csvData: CsvData,
  valueMap: ValueMap
): Promise<string> {
  const cleanedCsv = [csvData.orderedColumns.map((column) => column.id)];
  let currRow = 1;
  return new Promise((resolve, reject) => {
    Papa.parse(csvData.rawCsvFile || csvData.rawCsvUrl, {
      complete: () => {
        resolve(Papa.unparse(cleanedCsv));
      },
      step: (result) => {
        if (currRow >= csvData.firstDataRow && currRow <= csvData.lastDataRow) {
          const cleanedRow = result.data.map((cell) => {
            return cell in valueMap ? valueMap[cell] : cell;
          });
          cleanedCsv.push(cleanedRow);
        }
        currRow++;
      },
      error: () => {
        reject();
      },
    });
  });
}
