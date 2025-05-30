import { describe, it, expect, vi } from 'vitest';
import { adjustedData } from './utils.js';

describe('adjustedData', () => {
  vi.doMock('@/utils/constants', () => {
    return {
      baseURL: "unit-testing",
    };
  });
  it('should append the URL and process directory data correctly', () => {
    const input = {
      type: "directory",
      folders: [
        { name: "folder1", type: "directory" },
        { name: "folder2", type: "directory" },
      ],
      files: [
        { name: "file1.txt", type: "file" },
        { name: "file2.txt", type: "file" },
      ],
      path: "/root"
    };

    const url = "http://example.com/root/";

    const expected = {
      type: "directory",
      url: "http://example.com/root/",
      folders: [],
      files: [],
      items: [
        { name: "folder1", path: "/root/folder1", type: "directory", url: "http://example.com/root/folder1/" },
        { name: "folder2", path: "/root/folder2",  type: "directory", url: "http://example.com/root/folder2/" },
        { name: "file1.txt", path: "/root/file1.txt", type: "file", url: "http://example.com/root/file1.txt" },
        { name: "file2.txt", path: "/root/file2.txt", type: "file", url: "http://example.com/root/file2.txt" },
      ],
      path: "/root",
    };

    expect(adjustedData(input, url)).toEqual(expected);
  });

  it('should add a trailing slash to the URL if missing for a directory', () => {
    const input = { type: "directory", folders: [], files: [] };
    const url = "http://example.com/base";

    const expected = {
      type: "directory",
      url: "http://example.com/base/",
      folders: [],
      files: [],
      items: [],
    };

    expect(adjustedData(input, url)).toEqual(expected);
  });

  it('should handle non-directory types without modification to items', () => {
    const input = { type: "file", name: "file1.txt" };
    const url = "http://example.com/base";

    const expected = {
      type: "file",
      name: "file1.txt",
      url: "http://example.com/base",
    };

    expect(adjustedData(input, url)).toEqual(expected);
  });

  it('should handle missing folders and files gracefully', () => {
    const input = { type: "directory" };
    const url = "http://example.com/base";

    const expected = {
      type: "directory",
      url: "http://example.com/base/",
      items: [],
    };

    expect(adjustedData(input, url)).toEqual(expected);
  });

  it('should handle empty input object correctly', () => {
    const input = {};
    const url = "http://example.com/base";

    const expected = {
      url: "http://example.com/base",
    };

    expect(adjustedData(input, url)).toEqual(expected);
  });

});

